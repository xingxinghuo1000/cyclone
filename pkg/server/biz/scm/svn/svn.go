/*
Copyright 2017 caicloud authors. All rights reserved.

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

package svn

import (
	"github.com/caicloud/nirvana/log"

	c_v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	"github.com/caicloud/cyclone/pkg/server/apis/v1alpha1"
	"github.com/caicloud/cyclone/pkg/server/biz/scm"
	"github.com/caicloud/cyclone/pkg/util/cerr"
)

func init() {
	if err := scm.RegisterProvider(v1alpha1.SVN, NewSVN); err != nil {
		log.Errorln(err)
	}
}

// SVNUserPwdSep represents username and password separator, because SVN username can not contain ":".
const SVNUserPwdSep string = ":"

// SVN represents the SCM provider of SVN.
type SVN struct {
	scmCfg *v1alpha1.SCMSource
}

// NewSVN new SVN client
func NewSVN(scmCfg *v1alpha1.SCMSource) (scm.Provider, error) {
	return &SVN{scmCfg}, nil
}

// GetToken ...
func (s *SVN) GetToken() (string, error) {
	return s.scmCfg.Token, nil
}

// ListRepos ...
func (s *SVN) ListRepos() ([]scm.Repository, error) {
	return nil, cerr.ErrorNotImplemented.Error("list svn repos")
}

// ListBranches ...
func (s *SVN) ListBranches(repo string) ([]string, error) {
	return nil, cerr.ErrorNotImplemented.Error("list svn branches")
}

// ListTags ...
func (s *SVN) ListTags(repo string) ([]string, error) {
	return nil, cerr.ErrorNotImplemented.Error("list svn tags")
}

// ListDockerfiles ...
func (s *SVN) ListDockerfiles(repo string) ([]string, error) {
	return nil, cerr.ErrorNotImplemented.Error("list svn dockerfiles")
}

// CreateStatus ...
func (s *SVN) CreateStatus(status c_v1alpha1.StatusPhase, targetURL, repoURL, commitSha string) error {
	return cerr.ErrorNotImplemented.Error("create status")
}

// GetPullRequestSHA ...
func (s *SVN) GetPullRequestSHA(repoURL string, number int) (string, error) {
	return "", cerr.ErrorNotImplemented.Error("get pull request SHA")
}

// CheckToken ...
func (s *SVN) CheckToken() error {
	return nil
}

// CreateWebhook ...
func (s *SVN) CreateWebhook(repo string, webhook *scm.Webhook) error {
	return cerr.ErrorNotImplemented.Error("create svn webhook")
}

// DeleteWebhook ...
func (s *SVN) DeleteWebhook(repo string, webhookURL string) error {
	return cerr.ErrorNotImplemented.Error("delete svn webhook")
}
