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

package com.webank.wedatasphere.dss.appconn.mlss.project;

import com.webank.wedatasphere.dss.appconn.mlss.MLSSAppConn;
import com.webank.wedatasphere.dss.appconn.mlss.utils.MLSSConfig;
import com.webank.wedatasphere.dss.standard.app.sso.request.SSORequestOperation;
import com.webank.wedatasphere.dss.standard.app.structure.project.ProjectDeletionOperation;
import com.webank.wedatasphere.dss.standard.app.structure.project.ProjectService;
import com.webank.wedatasphere.dss.standard.app.structure.project.ProjectUpdateOperation;
import com.webank.wedatasphere.dss.standard.app.structure.project.ProjectUrlOperation;
import org.apache.linkis.httpclient.request.HttpAction;
import org.apache.linkis.httpclient.response.HttpResult;

import java.util.Map;

public class MLSSProjectService extends ProjectService {

    public MLSSProjectService(){
//        this.initMLSSConfig();
    }

    @Override
    public boolean isCooperationSupported() {
        return true;
    }

    @Override
    public boolean isProjectNameUnique() {
        return false;
    }

    @Override
    protected MLSSProjectCreationOperation createProjectCreationOperation() {
        SSORequestOperation<HttpAction, HttpResult> ssoRequestOperation = getSSORequestService().createSSORequestOperation(MLSSAppConn.MLSS_APPCONN_NAME);
        MLSSProjectCreationOperation MLSSProjectCreationOperation = new MLSSProjectCreationOperation(this, ssoRequestOperation);
        MLSSProjectCreationOperation.setStructureService(this);
        return MLSSProjectCreationOperation;
    }

    @Override
    protected ProjectUpdateOperation createProjectUpdateOperation() {
        return null;
    }

    @Override
    protected ProjectDeletionOperation createProjectDeletionOperation() {
        return null;
    }

    @Override
    protected ProjectUrlOperation createProjectUrlOperation() {
        return null;
    }


    protected void initMLSSConfig(){
        MLSSConfig.BASE_URL = this.getAppInstance().getBaseUrl();
        Map<String, Object> config = this.getAppInstance().getConfig();
        MLSSConfig.APP_KEY = String.valueOf(config.get("MLSS-SecretKey"));
        MLSSConfig.APP_SIGN = String.valueOf(config.get("MLSS-APPSignature"));
        MLSSConfig.AUTH_TYPE =  String.valueOf(config.get("MLSS-Auth-Type"));
        MLSSConfig.TIMESTAMP =  String.valueOf(config.get("MLSS-APPSignature"));
    }

}
