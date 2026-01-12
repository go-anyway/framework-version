// Copyright 2025 zampo.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// @contact  zampo3380@gmail.com

package version

import (
	"fmt"
	"runtime"
)

var (
	// Version 版本号，通过 -ldflags 注入
	Version = "dev"

	// CommitHash Git commit hash，通过 -ldflags 注入
	CommitHash = "unknown"

	// BuildTime 构建时间，通过 -ldflags 注入
	BuildTime = "unknown"

	// GoVersion Go 版本
	GoVersion = runtime.Version()
)

// Info 版本信息结构
type Info struct {
	Version    string
	CommitHash string
	BuildTime   string
	GoVersion  string
}

// GetInfo 获取版本信息
func GetInfo() Info {
	return Info{
		Version:    Version,
		CommitHash: CommitHash,
		BuildTime:  BuildTime,
		GoVersion:  GoVersion,
	}
}

// String 返回版本信息的字符串表示
func (i Info) String() string {
	return fmt.Sprintf("Version: %s\nCommit: %s\nBuildTime: %s\nGoVersion: %s",
		i.Version, i.CommitHash, i.BuildTime, i.GoVersion)
}

// ShortString 返回简短的版本信息
func (i Info) ShortString() string {
	return fmt.Sprintf("%s-%s", i.Version, i.CommitHash[:7])
}
