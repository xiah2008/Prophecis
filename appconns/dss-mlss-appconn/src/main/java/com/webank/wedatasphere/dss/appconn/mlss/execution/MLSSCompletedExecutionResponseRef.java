/*
 * Copyright 2019 WeBank
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package com.webank.wedatasphere.dss.appconn.mlss.execution;

import com.webank.wedatasphere.dss.standard.app.development.listener.common.CompletedExecutionResponseRef;
import com.webank.wedatasphere.dss.standard.app.development.listener.common.RefExecutionState;

import java.util.Map;

public class MLSSCompletedExecutionResponseRef extends CompletedExecutionResponseRef {

    public MLSSCompletedExecutionResponseRef(int status, String errorMessage){
        super(status);
        this.errorMsg = errorMessage;
    }

    public MLSSCompletedExecutionResponseRef(int status) {
        super(status);
    }

    public MLSSCompletedExecutionResponseRef(String responseBody, int status) {
        super(responseBody, status);
    }

    @Override
    public Map<String, Object> toMap() {
        return null;
    }

    @Override
    public String getErrorMsg() {
        return this.errorMsg;
    }

}
