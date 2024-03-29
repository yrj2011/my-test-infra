/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/glog"
)

// SecurityGroups: http://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.DescribeSecurityGroups
type SecurityGroups struct{}

type sgRef struct {
	id   string
	perm *ec2.IpPermission
}

func addRefs(refs map[string][]*sgRef, id string, acct string, perms []*ec2.IpPermission) {
	for _, perm := range perms {
		for _, pair := range perm.UserIdGroupPairs {
			// Ignore cross-account for now, and skip circular refs.
			if *pair.UserId == acct && *pair.GroupId != id {
				refs[*pair.GroupId] = append(refs[*pair.GroupId], &sgRef{id: id, perm: perm})
			}
		}
	}
}

func (SecurityGroups) MarkAndSweep(sess *session.Session, acct string, region string, set *Set) error {
	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})

	resp, err := svc.DescribeSecurityGroups(nil)
	if err != nil {
		return err
	}

	var toDelete []*securityGroup        // Deferred to disentangle referencing security groups
	ingress := make(map[string][]*sgRef) // sg.GroupId -> [sg.GroupIds with this ingress]
	egress := make(map[string][]*sgRef)  // sg.GroupId -> [sg.GroupIds with this egress]
	for _, sg := range resp.SecurityGroups {
		if *sg.GroupName == "default" {
			// TODO(zmerlynn): Is there really no better way to detect this?
			continue
		}
		s := &securityGroup{Account: acct, Region: region, ID: *sg.GroupId}
		addRefs(ingress, *sg.GroupId, acct, sg.IpPermissions)
		addRefs(egress, *sg.GroupId, acct, sg.IpPermissionsEgress)
		if set.Mark(s) {
			glog.Warningf("%s: deleting %T: %v", s.ARN(), sg, sg)
			toDelete = append(toDelete, s)
		}
	}
	for _, sg := range toDelete {
		for _, ref := range ingress[sg.ID] {
			glog.Infof("%v: revoking reference from %v", sg.ARN(), ref.id)
			_, err := svc.RevokeSecurityGroupIngress(&ec2.RevokeSecurityGroupIngressInput{
				GroupId:       aws.String(ref.id),
				IpPermissions: []*ec2.IpPermission{ref.perm},
			})
			if err != nil {
				glog.Warningf("%v: failed to revoke ingress reference from %v: %v", sg.ARN(), ref.id, err)
			}
		}
		for _, ref := range egress[sg.ID] {
			_, err := svc.RevokeSecurityGroupEgress(&ec2.RevokeSecurityGroupEgressInput{
				GroupId:       aws.String(ref.id),
				IpPermissions: []*ec2.IpPermission{ref.perm},
			})
			if err != nil {
				glog.Warningf("%v: failed to revoke egress reference from %v: %v", sg.ARN(), ref.id, err)
			}
		}
		_, err := svc.DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{GroupId: aws.String(sg.ID)})
		if err != nil {
			glog.Warningf("%v: delete failed: %v", sg.ARN(), err)
		}
	}
	return nil
}

type securityGroup struct {
	Account string
	Region  string
	ID      string
}

func (sg securityGroup) ARN() string {
	return fmt.Sprintf("arn:aws:ec2:%s:%s:security-group/%s", sg.Region, sg.Account, sg.ID)
}

func (sg securityGroup) ResourceKey() string {
	return sg.ARN()
}
