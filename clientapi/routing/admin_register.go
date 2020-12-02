// Copyright 2020 David Spenler
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

package routing

import (
	"net/http"

	"github.com/matrix-org/dendrite/userapi/api"

	"github.com/matrix-org/util"
)

type adminRegisterResponse struct {
	Test string `json:"test"`
}

// GetAdminRegister implements POST /admin/register
func GetAdminRegister(
	req *http.Request, userAPI api.UserInternalAPI,
	test string,
) util.JSONResponse {
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: adminRegisterResponse{
			Test: test,
		},
	}
}
