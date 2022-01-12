<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        {{tip}}容器配置配置
        <el-button size="small" style="margin-left:20px" @click="show=!show">
          {{!show?"展开":"收缩"}}
        </el-button>
        <el-button type="success" size="small" style="float: right;margin-right:50px" @click="addContainers" >添加{{tip}}Containers</el-button>
      </div>
    </template>
    {{form.Container}}
    <div v-show="show">
      <el-form :inline="true"  :model="form" ref="formRef"  >
        <div v-for="(item,index) in form.Container" class="mtb20">
          <div>
            <el-form-item
                label="容器名称"
                :rules="requireRules('容器名称填写')"
                :key="'Container_'+index+'.name'"
                :prop="'Container.'+index+'.name'"
            >
              <el-input v-model="item.name" placeholder="容器名称"/>
            </el-form-item>
            <el-form-item
                label="容器镜像"
                :rules="requireRules('容器镜像填写')"
                :key="'Container_'+index+'.image'"
                :prop="'Container.'+index+'.image'"
            >
              <el-input v-model="item.image" placeholder="容器镜像"/>
            </el-form-item>
            <el-form-item  >
              <el-button size="small" type="danger" @click="removeContainers(index)" >删除{{tip}}Containers</el-button>
            </el-form-item>
          </div>
          <el-divider></el-divider>
          <h3>添加网络端口 <el-button size="small" type="success"  class="mf20 mtb20" @click="addPort(index)" >添加端口</el-button></h3>
          <div v-for="(data,index2) in item.ports"  class="mtb20">
            <el-form-item
                label="port名称"
                :key="'Container_'+index+'.ports.'+index2+'.name'"
                :prop="'Container.'+index+'.ports.'+index2+'.name'"
            >
              <el-input v-model="data.name" placeholder="port名称"/>
            </el-form-item>
            <el-form-item
                label="containerPort"
                :key="'Container_'+index+'.ports.'+index2+'.containerPort'"
                :prop="'Container.'+index+'.ports.'+index2+'.containerPort'"
            >
              <el-input v-model="data.containerPort" placeholder="containerPort端口"/>
            </el-form-item>
            <el-form-item
                label="端口协议"
                :key="'Container_'+index+'.ports.'+index2+'.protocol'"
                :prop="'Container.'+index+'.ports.'+index2+'.protocol'"
            >
              <el-select  size="small" v-model="data.protocol"  class="select">
                <el-option label="TCP"  value="TCP"/>
                <el-option label="UDP"   value="UDP"/>
                <el-option label="UDP"   value="SCTP"/>
              </el-select>
            </el-form-item>
            <el-form-item
                label="hostPort"
                :key="'Container_'+index+'.ports.'+index2+'.hostPort'"
                :prop="'Container.'+index+'.ports.'+index2+'.hostPort'"
            >
              <el-input v-model="data.hostPort" placeholder="hostPort端口"/>
            </el-form-item>

            <el-form-item
                label="hostIP"
                :key="'Container_'+index+'.ports.'+index2+'.hostIP'"
                :prop="'Container.'+index+'.ports.'+index2+'.hostIP'"
            >
              <el-input v-model="data.hostIP" placeholder="hostIP"/>
            </el-form-item>
            <el-form-item  >
            <el-button size="small" type="danger" @click="removeItem(index,index2,'ports')" >删除端口</el-button>
          </el-form-item>
            <div class="dashed"></div>
          </div>
          <el-divider></el-divider>
          <div>
            <el-form-item
                label="容器执行命令"
                :key="'Container_'+index+'.command'"
                :prop="'Container.'+index+'.command'"
            >
              <el-input v-model="item.command"  class="longWidth" placeholder="容器执行命令"/>
            </el-form-item>
            <span  class="tip_info">多个值空格分割,并且空格使用双引号或者单引号包裹，例如 echo -c 'echo The app is running! && sleep 3600' </span>
          </div>
          <div>
            <el-form-item
                label="容器命令参数"
                :key="'Container_'+index+'.args'"
                :prop="'Container.'+index+'.args'"
            >
              <el-input v-model="item.args" class="longWidth" placeholder="容器执行命令参数"/>
            </el-form-item>
            <span  class="tip_info">多个值空格分割,并且空格使用双引号或者单引号包裹，例如  'The app is running! && sleep 3600' </span>
          </div>
          <div>
            <el-form-item
                label="工作目录"
                :key="'Container_'+index+'.workingDir'"
                :prop="'Container.'+index+'.workingDir'"
            >
              <el-input v-model="item.workingDir" class="longWidth" placeholder="工作目录"/>
            </el-form-item>
          </div>
          <el-divider></el-divider>
          <h2 class="advanced">高级应用-挂载相关
            <el-button size="small" style="margin-left:20px" @click="setAdvanced(0)">
            {{!advanced[0]?"展开":"收缩"}}
          </el-button></h2>
          <div v-show="advanced[0]"  class="advanced-box">
              <h3>环境变量相关-通过配置挂载 <el-button size="small" type="success"  class="mf20" @click="addEnvFrom(index)" >添加环境引用变量</el-button></h3>
              <div v-for="(envFromItem,index2) in item.envFrom"  class="mtb20">
                <el-form-item
                    label="前缀"
                    :key="'Container_'+index+'.envFrom.'+index2+'.prefix'"
                    :prop="'Container.'+index+'.envFrom.'+index2+'.prefix'"
                >
                  <el-input v-model="envFromItem.prefix"  placeholder="前缀"/>
                </el-form-item>
                <el-form-item  >
                  <el-button size="small" type="danger" @click="removeItem(index,index2,'envFrom')" >删除环境引用变量</el-button>
                </el-form-item>
                <div>
                  <el-form-item
                      label="配置名称"
                      :key="'Container_'+index+'.envFrom.'+index2+'.configMapRef.name'"
                      :prop="'Container.'+index+'.envFrom.'+index2+'.configMapRef.name'"
                  >

                    <el-select  size="small" v-model="envFromItem.configMapRef.name" class="select">
                      <el-option :key="item.name"  v-for="item in configMap" :label="item.name" :value="item.name"/>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="是否必须存在"
                                :key="'Container_'+index+'.envFrom.'+index2+'.configMapRef.optional'"
                                :prop="'Container.'+index+'.envFrom.'+index2+'.configMapRef.optional'"
                  >
                    <el-select  size="small" v-model="envFromItem.configMapRef.optional"  class="select">
                      <el-option label="false"  :value="false" />
                      <el-option label="true"   :value="true"  />
                    </el-select>
                  </el-form-item>
                </div>
                <div>
                  <el-form-item
                      label="密文名称"
                      :key="'Container_'+index+'.envFrom.'+index2+'.secretRef.name'"
                      :prop="'Container.'+index+'.envFrom.'+index2+'.secretRef.name'"
                  >

                    <el-select  size="small" v-model="envFromItem.secretRef.name" class="select">
                      <el-option :key="item.name"  v-for="item in secret" :label="item.name" :value="item.name"/>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="是否必须存在"
                                :key="'Container_'+index+'.envFrom.'+index2+'.secretRef.optional'"
                                :prop="'Container.'+index+'.envFrom.'+index2+'.secretRef.optional'"
                  >
                    <el-select  size="small" v-model="envFromItem.secretRef.optional"  class="select">
                      <el-option label="false"  :value="false" />
                      <el-option label="true"   :value="true"  />
                    </el-select>
                  </el-form-item>
                </div>
                <div class="dashed"></div>
              </div>
              <el-divider></el-divider>
              <h3>环境变量相关-通过设置挂载 <el-button size="small" type="success"  class="mf20" @click="addEnv(index)" >添加环境变量</el-button></h3>
              <div v-for="(envItem,index2) in item.env"  class="mtb20">
                <el-form-item
                    label="环境变量名"
                    :rules="requireRules('容器名称填写')"
                    :key="'Container_'+index+'.env.'+index2+'.name'"
                    :prop="'Container.'+index+'.env.'+index2+'.name'"
                >
                  <el-input v-model="envItem.name"  placeholder="变量名"/>
                </el-form-item>
                <el-form-item
                    label="环境变固定量值"
                    :key="'Container_'+index+'.env.'+index2+'.value'"
                    :prop="'Container.'+index+'.env.'+index2+'.value'"
                >
                  <el-input v-model="envItem.value"  placeholder="变量名"/>
                </el-form-item>
                <el-form-item  >
                  <el-button size="small" type="danger" @click="removeItem(index,index2,'env')" >删除环境变量</el-button>
                </el-form-item>
                <span  class="tip_info">如果设置固定值环境变量值和引用配置只能选择一个 </span>
                <div class="dashed"></div>
                <div v-if="envItem.value==''" class="mtb20">
                  <el-alert title="元数据对象一般用于将pod元数据挂载到容器环境变量" type="success" />
                  <h2 class="mtb20">子项设置元数据对象</h2>
                  <div>
                    <el-form-item
                        label="元数据对象-fieldRef.apiVersion"
                        :key="'Container.'+index+'.env.'+index2+'.valueFrom.fieldRef.apiVersion'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.fieldRef.apiVersion'"
                    >
                      <el-input v-model="envItem.valueFrom.fieldRef.apiVersion" placeholder="元数据对象-fieldRef.apiVersion" class="longWidth"/>
                    </el-form-item>
                    <el-form-item
                        label="元数据对象-fieldRef.FieldPath"
                        :key="'Container.'+index+'.env.'+index2+'.valueFrom.fieldRef.FieldPath'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.fieldRef.FieldPath'"
                    >
                      <el-input v-model="envItem.valueFrom.fieldRef.FieldPath" placeholder="元数据对象-fieldRef.FieldPath" class="longWidth" />
                    </el-form-item>
                  </div>
                  <h2  class="mtb20">子项设置资源对象</h2>
                  <div>
                    <el-form-item
                        label="资源对象-resourceFieldRef.containerName"
                        :key="'Container.'+index+'.env.'+index2+'.valueFrom.resourceFieldRef.containerName'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.resourceFieldRef.containerName'"
                    >
                      <el-input v-model="envItem.valueFrom.resourceFieldRef.containerName" placeholder="资源对象-resourceFieldRef.containerName" class="longWidth" />
                    </el-form-item>
                    <el-form-item
                        label="资源对象-resourceFieldRef.resource"
                        :key="'Container.'+index+'.env.'+index2+'.valueFrom.resourceFieldRef.resource'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.resourceFieldRef.resource'"
                    >
                      <el-input v-model="envItem.valueFrom.resourceFieldRef.resource" placeholder="资源对象-resourceFieldRef.resource" class="longWidth"/>
                    </el-form-item>
                    <el-form-item
                        label="资源对象单位"
                        :key="'Container.'+index+'.env.'+index2+'.valueFrom.resourceFieldRef.divisor'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.resourceFieldRef.divisor'"
                    >
                      <el-input v-model="envItem.valueFrom.resourceFieldRef.divisor" placeholder="资源对象-resourceFieldRef.divisor" />
                    </el-form-item>
                  </div>
                  <el-divider></el-divider>
                  <el-alert title="通过其他资源挂载环境变量" type="warning" />
                  <div class="mtb20">
                    <el-form-item
                        label="配置名称"
                        :key="'Container_'+index+'.env.'+index2+'.valueFrom.configMapKeyRef.name'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.configMapKeyRef.name'"
                    >

                      <el-select  size="small" v-model="envItem.valueFrom.configMapKeyRef.name" class="select">
                        <el-option :key="item.name"  v-for="item in configMap" :label="item.name" :value="item.name"/>
                      </el-select>
                    </el-form-item>
                    <el-form-item
                        label="key"
                        :key="'Container.'+index+'.env.'+index2+'.valueFrom.configMapKeyRef.key'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.configMapKeyRef.key'"
                    >
                      <el-input v-model="envItem.valueFrom.configMapKeyRef.key" placeholder="key" class="longWidth" />
                    </el-form-item>
                    <el-form-item label="是否必须存在"
                                  :key="'Container_'+index+'.env.'+index2+'.valueFrom.configMapKeyRef.optional'"
                                  :prop="'Container.'+index+'.env.'+index2+'.valueFrom.configMapKeyRef.optional'"
                    >
                      <el-select  size="small" v-model="envItem.valueFrom.configMapKeyRef.optional"  class="select">
                        <el-option label="false"  :value="false" />
                        <el-option label="true"   :value="true"  />
                      </el-select>
                    </el-form-item>
                  </div>
                  <div>
                    <el-form-item
                        label="密文名称"
                        :key="'Container_'+index+'.env.'+index2+'.valueFrom.secretKeyRef.name'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.secretKeyRef.name'"
                    >

                      <el-select  size="small" v-model="envItem.valueFrom.secretKeyRef.name" class="select">
                        <el-option :key="item.name"  v-for="item in secret" :label="item.name" :value="item.name"/>
                      </el-select>
                    </el-form-item>
                    <el-form-item
                        label="key"
                        :key="'Container.'+index+'.env.'+index2+'.valueFrom.secretKeyRef.key'"
                        :prop="'Container.'+index+'.env.'+index2+'.valueFrom.secretKeyRef.key'"
                    >
                      <el-input v-model="envItem.valueFrom.secretKeyRef.key" placeholder="key" class="longWidth" />
                    </el-form-item>
                    <el-form-item label="是否必须存在"
                                  :key="'Container_'+index+'.env.'+index2+'.valueFrom.secretKeyRef.optional'"
                                  :prop="'Container.'+index+'.env.'+index2+'.valueFrom.secretKeyRef.optional'"
                    >
                      <el-select  size="small" v-model="envItem.valueFrom.secretKeyRef.optional"  class="select">
                        <el-option label="false"  :value="false" />
                        <el-option label="true"   :value="true"  />
                      </el-select>
                    </el-form-item>
                  </div>
                  <div class="dashed"></div>
                </div>
              </div>
              <el-divider></el-divider>
              <h3 >数据卷挂载<el-button size="small" type="success"  class="mf20" @click="addVolumeMounts(index)" >添加数据卷挂载</el-button></h3>
                <div v-for="(volumeMount,index2) in item.volumeMounts"  class="mtb20">
                  <div>
                    <el-form-item
                      label="挂载名称"
                      :rules="requireRules('挂载名称必须填写')"
                      :key="'Container_'+index+'.volumeMounts.'+index2+'.name'"
                      :prop="'Container.'+index+'.volumeMounts.'+index2+'.name'"
                  >
                    <el-input v-model="volumeMount.name"  placeholder="挂载名称"/>
                  </el-form-item>
                  <el-form-item
                      label="挂载路径"
                      :rules="requireRules('挂载名称必须填写')"
                      :key="'Container_'+index+'.volumeMounts.'+index2+'.mountPath'"
                      :prop="'Container.'+index+'.volumeMounts.'+index2+'.mountPath'"
                  >
                    <el-input v-model="volumeMount.mountPath"  placeholder="挂载路径"/>
                  </el-form-item>
                  <el-form-item
                      label="是否只读"
                      :key="'Container_'+index+'.volumeMounts.'+index2+'.readOnly'"
                      :prop="'Container.'+index+'.volumeMounts.'+index2+'.readOnly'"
                  >
                    <el-select  size="small" v-model="volumeMount.readOnly" class="select">
                      <el-option :label="true" :value="true" />
                      <el-option :label="false" :value="false" />
                    </el-select>
                  </el-form-item>
                    <el-form-item  >
                      <el-button size="small" type="danger" @click="removeItem(index,index2,'volumeMounts')" >删除数据卷挂载</el-button>
                    </el-form-item>
                  </div>
                  <div>
                    <el-form-item
                        label="挂载子路径"
                        :key="'Container_'+index+'.volumeMounts.'+index2+'.subPath'"
                        :prop="'Container.'+index+'.volumeMounts.'+index2+'.subPath'"
                    >
                      <el-input v-model="volumeMount.subPath"  placeholder="挂载子路径"/>
                    </el-form-item>
                    <el-form-item
                        v-if="volumeMount.subPath==''"
                        label="挂载子路径表达式"
                        :key="'Container_'+index+'.volumeMounts.'+index2+'.subPathExpr'"
                        :prop="'Container.'+index+'.volumeMounts.'+index2+'.subPathExpr'"
                    >
                      <el-input v-model="volumeMount.subPathExpr"  placeholder="挂载子路径表达式"/>
                    </el-form-item>
                    <el-form-item
                        label="mountPropagation"
                        :key="'Container_'+index+'.volumeMounts.'+index2+'.mountPropagation'"
                        :prop="'Container.'+index+'.volumeMounts.'+index2+'.mountPropagation'"
                    >
                      <el-select  size="small" v-model="volumeMount.mountPropagation" class="select">
                        <el-option label="None" value="None" />
                        <el-option label="HostToContainer" value="HostToContainer" />
                        <el-option label="Bidirectional" value="Bidirectional" />
                      </el-select>
                    </el-form-item>

                  </div>
                  <div class="dashed"></div>
                </div>
              <el-divider></el-divider>
          </div>
          <h2 class="advanced">高级应用-限制相关
            <el-button size="small" style="margin-left:20px" @click="setAdvanced(1)">
              {{!advanced[1]?"展开":"收缩"}}
            </el-button></h2>
          <div v-show="advanced[1]"  class="advanced-box">
            <h3 >资源限制<el-button size="small" class="mf20" type="success"  @click="setResources(index)" >添加资源限制</el-button></h3>
            <div v-if="item.resources.limits"  class="mtb20">
              <div>
                <el-form-item
                    label="cpu最大限制"
                    :key="'Container.'+index+'.resources.limits.cpu'"
                    :prop="'Container.'+index+'.resources.limits.cpu'"
                >
                  <el-input v-model="item.resources.limits.cpu" placeholder="cpu限制" />
                </el-form-item>
                <el-form-item
                    label="内存最大限制"
                    :key="'Container.'+index+'.resources.limits.memory'"
                    :prop="'Container.'+index+'.resources.limits.memory'"
                >
                  <el-input v-model="item.resources.limits.memory" placeholder="memory限制" />
                </el-form-item>
              </div>
              <div>
                <el-form-item
                    label="申请cpu"
                    :key="'Container.'+index+'.resources.requests.cpu'"
                    :prop="'Container.'+index+'.resources.requests.cpu'"
                >
                  <el-input v-model="item.resources.requests.cpu" placeholder="cpu限制" />
                </el-form-item>
                <el-form-item
                    label="申请内存"
                    :key="'Container.'+index+'.resources.requests.memory'"
                    :prop="'Container.'+index+'.resources.requests.memory'"
                >
                  <el-input v-model="item.resources.requests.memory" placeholder="memory限制" />
                </el-form-item>
                <el-form-item  >
                  <el-button size="small" type="danger" @click="removeResources(index,'resources')" >取消资源限制</el-button>
                </el-form-item>
              </div>
            </div>
            <el-divider></el-divider>
          </div>
          <h2 class="advanced">高级应用-探针-生命钩子相关
            <el-button size="small" style="margin-left:20px" @click="setAdvanced(2)">
              {{!advanced[2]?"展开":"收缩"}}
            </el-button></h2>
          <div v-show="advanced[2]"  class="advanced-box">
              <h3 >存活探针<el-button size="small" type="success" class="mf20" @click="setLivenessProbe(index)" >添加存活探针</el-button></h3>
              <div v-if="item.livenessProbe&&item.livenessProbe.type"  class="mtb20">
                <div>

                  <el-form-item
                      label="initialDelaySeconds"
                      :key="'Container.'+index+'.livenessProbe.initialDelaySeconds'"
                      :prop="'Container.'+index+'.livenessProbe.initialDelaySeconds'"
                  >
                    <el-input v-model="item.livenessProbe.initialDelaySeconds" placeholder="initialDelaySeconds" />
                  </el-form-item>
                  <el-form-item
                      label="timeoutSeconds"
                      :key="'Container.'+index+'.livenessProbe.timeoutSeconds'"
                      :prop="'Container.'+index+'.livenessProbe.timeoutSeconds'"
                  >
                    <el-input v-model="item.livenessProbe.timeoutSeconds" placeholder="timeoutSeconds" />
                  </el-form-item>
                  <el-form-item
                      label="periodSeconds"
                      :key="'Container.'+index+'.livenessProbe.periodSeconds'"
                      :prop="'Container.'+index+'.livenessProbe.periodSeconds'"
                  >
                    <el-input v-model="item.livenessProbe.periodSeconds" placeholder="periodSeconds" />
                  </el-form-item>
                </div>
                <div>
                  <el-form-item
                      label="successThreshold"
                      :key="'Container.'+index+'.livenessProbe.successThreshold'"
                      :prop="'Container.'+index+'.livenessProbe.successThreshold'"
                  >
                    <el-input v-model="item.livenessProbe.successThreshold" placeholder="successThreshold" />
                  </el-form-item>
                  <el-form-item
                      label="failureThreshold"
                      :key="'Container.'+index+'.livenessProbe.failureThreshold'"
                      :prop="'Container.'+index+'.livenessProbe.failureThreshold'"
                  >
                    <el-input v-model="item.livenessProbe.successThreshold" placeholder="failureThreshold" />
                  </el-form-item>
                  <el-form-item
                      label="terminationGracePeriodSeconds"
                      :key="'Container.'+index+'.livenessProbe.terminationGracePeriodSeconds'"
                      :prop="'Container.'+index+'.livenessProbe.terminationGracePeriodSeconds'"
                  >
                    <el-input v-model="item.livenessProbe.successThreshold" placeholder="terminationGracePeriodSeconds" />
                  </el-form-item>
                  <el-form-item  >
                    <el-button size="small" type="danger" @click="removeResources(index,'livenessProbe')" >取消存活探针</el-button>
                  </el-form-item>
                </div>
                <div>
                  <el-form-item label="设置探针类型"
                                :key="'Container.'+index+'.livenessProbe.type'"
                                :prop="'Container.'+index+'.livenessProbe.type'"
                  >
                    <el-select  size="small" v-model="item.livenessProbe.type"  class="select">
                      <el-option label="exec"  value="exec" />
                      <el-option label="httpGet"   value="httpGet"  />
                      <el-option label="tcpSocket"   value="tcpSocket"  />
                    </el-select>
                  </el-form-item>
                  <div v-if="item.livenessProbe.type=='exec'">
                    <el-form-item
                        label="执行命令"
                        :key="'Container.'+index+'.livenessProbe.exec.command'"
                        :prop="'Container.'+index+'.livenessProbe.exec.command'"
                    >
                      <el-input v-model="item.livenessProbe.exec.command"  class="longWidth" placeholder="执行命令"/>
                    </el-form-item>
                    <span  class="tip_info">多个值空格分割,并且空格使用双引号或者单引号包裹，例如 echo -c 'echo The app is running! && sleep 3600' </span>
                  </div>
                  <div v-if="item.livenessProbe.type=='httpGet'">
                    <el-form-item
                        label="Host"
                        :key="'Container.'+index+'.livenessProbe.httpGet.host'"
                        :prop="'Container.'+index+'.livenessProbe.httpGet.host'"
                    >
                      <el-input v-model="item.livenessProbe.httpGet.host"  class="longWidth" placeholder="host"/>
                    </el-form-item>
                    <el-form-item
                        label="端口"
                        :key="'Container.'+index+'.livenessProbe.httpGet.port'"
                        :prop="'Container.'+index+'.livenessProbe.httpGet.port'"
                    >
                      <el-input v-model="item.livenessProbe.httpGet.port"  class="longWidth" placeholder="port"/>
                    </el-form-item>
                    <el-form-item
                        label="路径"
                        :key="'Container.'+index+'.livenessProbe.httpGet.path'"
                        :prop="'Container.'+index+'.livenessProbe.httpGet.path'"
                    >
                      <el-input v-model="item.livenessProbe.httpGet.path"  class="longWidth" placeholder="路径"/>
                    </el-form-item>
                    <el-form-item label="协议类型"
                                  :key="'Container.'+index+'.livenessProbe.httpGet.scheme'"
                                  :prop="'Container.'+index+'.livenessProbe.httpGet.scheme'"
                    >
                      <el-select  size="small" v-model="item.livenessProbe.httpGet.scheme"  class="select">
                        <el-option label="https"  value="https" />
                        <el-option label="http"   value="http"  />
                      </el-select>
                    </el-form-item>
                    <el-form-item  >
                      <el-button size="small" type="success" @click="addHttpHeaders(index,'livenessProbe')" >添加header</el-button>
                    </el-form-item>
                    <div v-for="(data,index2) in item.livenessProbe.httpGet.httpHeaders"  class="mtb20">
                      <el-form-item
                          label="header名称"
                          :key="'Container.'+index+'.livenessProbe.httpGet.httpHeaders.'+index2+'.name'"
                          :prop="'Container.'+index+'.livenessProbe.httpGet.httpHeaders.'+index2+'.name'"
                      >
                        <el-input v-model="data.name" placeholder="header名称"/>
                      </el-form-item>
                      <el-form-item
                          label="header值"
                          :key="'Container.'+index+'.livenessProbe.httpGet.httpHeaders.'+index2+'.value'"
                          :prop="'Container.'+index+'.livenessProbe.httpGet.httpHeaders.'+index2+'.value'"
                      >
                        <el-input v-model="data.value" placeholder="header值"/>
                      </el-form-item>
                      <el-form-item  >
                        <el-button size="small" type="danger" @click="removeHeader(index,index2,'livenessProbe')" >删除</el-button>
                      </el-form-item>
                    </div>
                  </div>
                  <div v-if="item.livenessProbe.type=='tcpSocket'">
                    <el-form-item
                        label="host"
                        :key="'Container.'+index+'.livenessProbe.tcpSocket.host'"
                        :prop="'Container.'+index+'.livenessProbe.tcpSocket.host'"
                    >
                      <el-input v-model="item.livenessProbe.tcpSocket.host"  class="longWidth" placeholder="主机名称"/>
                    </el-form-item>
                    <el-form-item
                        label="主机端口"
                        :key="'Container.'+index+'.livenessProbe.tcpSocket.port'"
                        :prop="'Container.'+index+'.livenessProbe.tcpSocket.port'"
                    >
                      <el-input v-model="item.livenessProbe.tcpSocket.port"  class="longWidth" placeholder="主机port"/>
                    </el-form-item>
                  </div>
                </div>
              </div>
              <div class="bar"></div>
              <h3 >就绪探针<el-button size="small" type="success"  class="mf20" @click="setReadinessProbe(index)" >添加存活探针</el-button></h3>
              <div v-if="item.readinessProbe&&item.readinessProbe.type"  class="mtb20">
                <div>

                  <el-form-item
                      label="initialDelaySeconds"
                      :key="'Container.'+index+'.readinessProbe.initialDelaySeconds'"
                      :prop="'Container.'+index+'.readinessProbe.initialDelaySeconds'"
                  >
                    <el-input v-model="item.readinessProbe.initialDelaySeconds" placeholder="initialDelaySeconds" />
                  </el-form-item>
                  <el-form-item
                      label="timeoutSeconds"
                      :key="'Container.'+index+'.readinessProbe.timeoutSeconds'"
                      :prop="'Container.'+index+'.readinessProbe.timeoutSeconds'"
                  >
                    <el-input v-model="item.readinessProbe.timeoutSeconds" placeholder="timeoutSeconds" />
                  </el-form-item>
                  <el-form-item
                      label="periodSeconds"
                      :key="'Container.'+index+'.readinessProbe.periodSeconds'"
                      :prop="'Container.'+index+'.readinessProbe.periodSeconds'"
                  >
                    <el-input v-model="item.readinessProbe.periodSeconds" placeholder="periodSeconds" />
                  </el-form-item>
                </div>
                <div>
                  <el-form-item
                      label="successThreshold"
                      :key="'Container.'+index+'.readinessProbe.successThreshold'"
                      :prop="'Container.'+index+'.readinessProbe.successThreshold'"
                  >
                    <el-input v-model="item.readinessProbe.successThreshold" placeholder="successThreshold" />
                  </el-form-item>
                  <el-form-item
                      label="failureThreshold"
                      :key="'Container.'+index+'.readinessProbe.failureThreshold'"
                      :prop="'Container.'+index+'.readinessProbe.failureThreshold'"
                  >
                    <el-input v-model="item.readinessProbe.successThreshold" placeholder="failureThreshold" />
                  </el-form-item>
                  <el-form-item
                      label="terminationGracePeriodSeconds"
                      :key="'Container.'+index+'.readinessProbe.terminationGracePeriodSeconds'"
                      :prop="'Container.'+index+'.readinessProbe.terminationGracePeriodSeconds'"
                  >
                    <el-input v-model="item.readinessProbe.successThreshold" placeholder="terminationGracePeriodSeconds" />
                  </el-form-item>
                  <el-form-item  >
                    <el-button size="small" type="danger" @click="removeResources(index,'readinessProbe')" >取消就绪探针</el-button>
                  </el-form-item>
                </div>
                <div>
                  <el-form-item label="设置探针类型"
                                :key="'Container.'+index+'.readinessProbe.type'"
                                :prop="'Container.'+index+'.readinessProbe.type'"
                  >
                    <el-select  size="small" v-model="item.readinessProbe.type"  class="select">
                      <el-option label="exec"  value="exec" />
                      <el-option label="httpGet"   value="httpGet"  />
                      <el-option label="tcpSocket"   value="tcpSocket"  />
                    </el-select>
                  </el-form-item>
                  <div v-if="item.readinessProbe.type=='exec'">
                    <el-form-item
                        label="执行命令"
                        :key="'Container.'+index+'.readinessProbe.exec.command'"
                        :prop="'Container.'+index+'.readinessProbe.exec.command'"
                    >
                      <el-input v-model="item.readinessProbe.exec.command"  class="longWidth" placeholder="执行命令"/>
                    </el-form-item>
                    <span  class="tip_info">多个值空格分割,并且空格使用双引号或者单引号包裹，例如 echo -c 'echo The app is running! && sleep 3600' </span>
                  </div>
                  <div v-if="item.readinessProbe.type=='httpGet'">
                    <el-form-item
                        label="Host"
                        :key="'Container.'+index+'.readinessProbe.httpGet.host'"
                        :prop="'Container.'+index+'.readinessProbe.httpGet.host'"
                    >
                      <el-input v-model="item.readinessProbe.httpGet.host"  class="longWidth" placeholder="host"/>
                    </el-form-item>
                    <el-form-item
                        label="端口"
                        :key="'Container.'+index+'.readinessProbe.httpGet.port'"
                        :prop="'Container.'+index+'.readinessProbe.httpGet.port'"
                    >
                      <el-input v-model="item.readinessProbe.httpGet.port"  class="longWidth" placeholder="port"/>
                    </el-form-item>
                    <el-form-item
                        label="路径"
                        :key="'Container.'+index+'.readinessProbe.httpGet.path'"
                        :prop="'Container.'+index+'.readinessProbe.httpGet.path'"
                    >
                      <el-input v-model="item.readinessProbe.httpGet.path"  class="longWidth" placeholder="路径"/>
                    </el-form-item>
                    <el-form-item label="协议类型"
                                  :key="'Container.'+index+'.readinessProbe.httpGet.scheme'"
                                  :prop="'Container.'+index+'.readinessProbe.httpGet.scheme'"
                    >
                      <el-select  size="small" v-model="item.readinessProbe.httpGet.scheme"  class="select">
                        <el-option label="https"  value="https" />
                        <el-option label="http"   value="http"  />
                      </el-select>
                    </el-form-item>
                    <el-form-item  >
                      <el-button size="small" type="success" @click="addHttpHeaders(index,'readinessProbe')" >添加header</el-button>
                    </el-form-item>
                    <div v-for="(data,index2) in item.readinessProbe.httpGet.httpHeaders"  class="mtb20">
                      <el-form-item
                          label="header名称"
                          :key="'Container.'+index+'.readinessProbe.httpGet.httpHeaders.'+index2+'.name'"
                          :prop="'Container.'+index+'.readinessProbe.httpGet.httpHeaders.'+index2+'.name'"
                      >
                        <el-input v-model="data.name" placeholder="header名称"/>
                      </el-form-item>
                      <el-form-item
                          label="header值"
                          :key="'Container.'+index+'.readinessProbe.httpGet.httpHeaders.'+index2+'.value'"
                          :prop="'Container.'+index+'.readinessProbe.httpGet.httpHeaders.'+index2+'.value'"
                      >
                        <el-input v-model="data.value" placeholder="header值"/>
                      </el-form-item>
                      <el-form-item  >
                        <el-button size="small" type="danger" @click="removeHeader(index,index2,'readinessProbe')" >删除</el-button>
                      </el-form-item>
                    </div>
                  </div>
                  <div v-if="item.readinessProbe.type=='tcpSocket'">
                    <el-form-item
                        label="host"
                        :key="'Container.'+index+'.readinessProbe.tcpSocket.host'"
                        :prop="'Container.'+index+'.readinessProbe.tcpSocket.host'"
                    >
                      <el-input v-model="item.readinessProbe.tcpSocket.host"  class="longWidth" placeholder="主机名称"/>
                    </el-form-item>
                    <el-form-item
                        label="主机端口"
                        :key="'Container.'+index+'.readinessProbe.tcpSocket.port'"
                        :prop="'Container.'+index+'.readinessProbe.tcpSocket.port'"
                    >
                      <el-input v-model="item.readinessProbe.tcpSocket.port"  class="longWidth" placeholder="主机port"/>
                    </el-form-item>
                  </div>
                </div>
              </div>
              <div class="bar"></div>
              <h3 >启动探针<el-button size="small" type="success"  class="mf20" @click="setStartupProbe(index)" >添加存活探针</el-button></h3>
              <div v-if="item.startupProbe&&item.startupProbe.type"  class="mtb20">
                <div>

                  <el-form-item
                      label="initialDelaySeconds"
                      :key="'Container.'+index+'.startupProbe.initialDelaySeconds'"
                      :prop="'Container.'+index+'.startupProbe.initialDelaySeconds'"
                  >
                    <el-input v-model="item.startupProbe.initialDelaySeconds" placeholder="initialDelaySeconds" />
                  </el-form-item>
                  <el-form-item
                      label="timeoutSeconds"
                      :key="'Container.'+index+'.startupProbe.timeoutSeconds'"
                      :prop="'Container.'+index+'.startupProbe.timeoutSeconds'"
                  >
                    <el-input v-model="item.startupProbe.timeoutSeconds" placeholder="timeoutSeconds" />
                  </el-form-item>
                  <el-form-item
                      label="periodSeconds"
                      :key="'Container.'+index+'.startupProbe.periodSeconds'"
                      :prop="'Container.'+index+'.startupProbe.periodSeconds'"
                  >
                    <el-input v-model="item.startupProbe.periodSeconds" placeholder="periodSeconds" />
                  </el-form-item>
                </div>
                <div>
                  <el-form-item
                      label="successThreshold"
                      :key="'Container.'+index+'.startupProbe.successThreshold'"
                      :prop="'Container.'+index+'.startupProbe.successThreshold'"
                  >
                    <el-input v-model="item.startupProbe.successThreshold" placeholder="successThreshold" />
                  </el-form-item>
                  <el-form-item
                      label="failureThreshold"
                      :key="'Container.'+index+'.startupProbe.failureThreshold'"
                      :prop="'Container.'+index+'.startupProbe.failureThreshold'"
                  >
                    <el-input v-model="item.startupProbe.successThreshold" placeholder="failureThreshold" />
                  </el-form-item>
                  <el-form-item
                      label="terminationGracePeriodSeconds"
                      :key="'Container.'+index+'.startupProbe.terminationGracePeriodSeconds'"
                      :prop="'Container.'+index+'.startupProbe.terminationGracePeriodSeconds'"
                  >
                    <el-input v-model="item.startupProbe.successThreshold" placeholder="terminationGracePeriodSeconds" />
                  </el-form-item>
                  <el-form-item  >
                    <el-button size="small" type="danger" @click="removeResources(index,'startupProbe')" >取消启动探针</el-button>
                  </el-form-item>
                </div>
                <div>
                  <el-form-item label="设置探针类型"
                                :key="'Container.'+index+'.startupProbe.type'"
                                :prop="'Container.'+index+'.startupProbe.type'"
                  >
                    <el-select  size="small" v-model="item.startupProbe.type"  class="select">
                      <el-option label="exec"  value="exec" />
                      <el-option label="httpGet"   value="httpGet"  />
                      <el-option label="tcpSocket"   value="tcpSocket"  />
                    </el-select>
                  </el-form-item>
                  <div v-if="item.startupProbe.type=='exec'">
                    <el-form-item
                        label="执行命令"
                        :key="'Container.'+index+'.startupProbe.exec.command'"
                        :prop="'Container.'+index+'.startupProbe.exec.command'"
                    >
                      <el-input v-model="item.startupProbe.exec.command"  class="longWidth" placeholder="执行命令"/>
                    </el-form-item>
                    <span  class="tip_info">多个值空格分割,并且空格使用双引号或者单引号包裹，例如 echo -c 'echo The app is running! && sleep 3600' </span>
                  </div>
                  <div v-if="item.startupProbe.type=='httpGet'">
                    <el-form-item
                        label="Host"
                        :key="'Container.'+index+'.startupProbe.httpGet.host'"
                        :prop="'Container.'+index+'.startupProbe.httpGet.host'"
                    >
                      <el-input v-model="item.startupProbe.httpGet.host"  class="longWidth" placeholder="host"/>
                    </el-form-item>
                    <el-form-item
                        label="端口"
                        :key="'Container.'+index+'.startupProbe.httpGet.port'"
                        :prop="'Container.'+index+'.startupProbe.httpGet.port'"
                    >
                      <el-input v-model="item.startupProbe.httpGet.port"  class="longWidth" placeholder="port"/>
                    </el-form-item>
                    <el-form-item
                        label="路径"
                        :key="'Container.'+index+'.startupProbe.httpGet.path'"
                        :prop="'Container.'+index+'.startupProbe.httpGet.path'"
                    >
                      <el-input v-model="item.startupProbe.httpGet.path"  class="longWidth" placeholder="路径"/>
                    </el-form-item>
                    <el-form-item label="协议类型"
                                  :key="'Container.'+index+'.startupProbe.httpGet.scheme'"
                                  :prop="'Container.'+index+'.startupProbe.httpGet.scheme'"
                    >
                      <el-select  size="small" v-model="item.startupProbe.httpGet.scheme"  class="select">
                        <el-option label="https"  value="https" />
                        <el-option label="http"   value="http"  />
                      </el-select>
                    </el-form-item>
                    <el-form-item  >
                      <el-button size="small" type="success" @click="addHttpHeaders(index,'startupProbe')" >添加header</el-button>
                    </el-form-item>
                    <div v-for="(data,index2) in item.startupProbe.httpGet.httpHeaders"  class="mtb20">
                      <el-form-item
                          label="header名称"
                          :key="'Container.'+index+'.startupProbe.httpGet.httpHeaders.'+index2+'.name'"
                          :prop="'Container.'+index+'.startupProbe.httpGet.httpHeaders.'+index2+'.name'"
                      >
                        <el-input v-model="data.name" placeholder="header名称"/>
                      </el-form-item>
                      <el-form-item
                          label="header值"
                          :key="'Container.'+index+'.startupProbe.httpGet.httpHeaders.'+index2+'.value'"
                          :prop="'Container.'+index+'.startupProbe.httpGet.httpHeaders.'+index2+'.value'"
                      >
                        <el-input v-model="data.value" placeholder="header值"/>
                      </el-form-item>
                      <el-form-item  >
                        <el-button size="small" type="danger" @click="removeHeader(index,index2,'startupProbe')" >删除</el-button>
                      </el-form-item>
                    </div>
                  </div>
                  <div v-if="item.startupProbe.type=='tcpSocket'">
                    <el-form-item
                        label="host"
                        :key="'Container.'+index+'.startupProbe.tcpSocket.host'"
                        :prop="'Container.'+index+'.startupProbe.tcpSocket.host'"
                    >
                      <el-input v-model="item.startupProbe.tcpSocket.host"  class="longWidth" placeholder="主机名称"/>
                    </el-form-item>
                    <el-form-item
                        label="主机端口"
                        :key="'Container.'+index+'.startupProbe.tcpSocket.port'"
                        :prop="'Container.'+index+'.startupProbe.tcpSocket.port'"
                    >
                      <el-input v-model="item.startupProbe.tcpSocket.port"  class="longWidth" placeholder="主机port"/>
                    </el-form-item>
                  </div>
                </div>
              </div>
              <div class="bar"></div>
              <el-divider></el-divider>
              <h3 >生命周期钩子postStart<el-button size="small" type="success"  class="mf20" @click="setResourcesPostStart(index)" >添加postStart</el-button></h3>
              <div v-if="item.lifecycle&&item.lifecycle.postStart" class="mtb20">
                <el-form-item label="设置探针类型"
                              :key="'Container.'+index+'.lifecycle.postStart.type'"
                              :prop="'Container.'+index+'.lifecycle.postStart.type'"
                >
                  <el-select  size="small" v-model="item.lifecycle.postStart.type"  class="select">
                    <el-option label="exec"  value="exec" />
                    <el-option label="httpGet"   value="httpGet"  />
                    <el-option label="tcpSocket"   value="tcpSocket"  />
                  </el-select>
                </el-form-item>
                <el-form-item  >
                  <el-button size="small" type="danger" @click="removeResources(index,'lifecycle','postStart')" >取消postStart</el-button>
                </el-form-item>
                <div v-if="item.lifecycle.postStart.type=='exec'">
                  <el-form-item
                      label="执行命令"
                      :key="'Container.'+index+'.lifecycle.postStart.exec.command'"
                      :prop="'Container.'+index+'.lifecycle.postStart.exec.command'"
                  >
                    <el-input v-model="item.lifecycle.postStart.exec.command"  class="longWidth" placeholder="执行命令"/>
                  </el-form-item>
                  <span  class="tip_info">多个值空格分割,并且空格使用双引号或者单引号包裹，例如 echo -c 'echo The app is running! && sleep 3600' </span>
                </div>
                <div v-if="item.lifecycle.postStart.type=='httpGet'">
                  <el-form-item
                      label="Host"
                      :key="'Container.'+index+'.lifecycle.postStart.httpGet.host'"
                      :prop="'Container.'+index+'.lifecycle.postStart.httpGet.host'"
                  >
                    <el-input v-model="item.lifecycle.postStart.httpGet.host"  class="longWidth" placeholder="host"/>
                  </el-form-item>
                  <el-form-item
                      label="端口"
                      :key="'Container.'+index+'.lifecycle.postStart.httpGet.port'"
                      :prop="'Container.'+index+'.lifecycle.postStart.httpGet.port'"
                  >
                    <el-input v-model="item.lifecycle.postStart.httpGet.port"  class="longWidth" placeholder="port"/>
                  </el-form-item>
                  <el-form-item
                      label="路径"
                      :key="'Container.'+index+'.lifecycle.postStart.httpGet.path'"
                      :prop="'Container.'+index+'.lifecycle.postStart.httpGet.path'"
                  >
                    <el-input v-model="item.lifecycle.postStart.httpGet.path"  class="longWidth" placeholder="路径"/>
                  </el-form-item>
                  <el-form-item label="协议类型"
                                :key="'Container.'+index+'.lifecycle.postStart.httpGet.scheme'"
                                :prop="'Container.'+index+'.lifecycle.postStart.httpGet.scheme'"
                  >
                    <el-select  size="small" v-model="item.lifecycle.postStart.httpGet.scheme"  class="select">
                      <el-option label="https"  value="https" />
                      <el-option label="http"   value="http"  />
                    </el-select>
                  </el-form-item>
                  <el-form-item  >
                    <el-button size="small" type="success" @click="addHttpHeaders(index,'lifecycle.postStart')" >添加header</el-button>
                  </el-form-item>
                  <div v-for="(data,index2) in item.lifecycle.postStart.httpGet.httpHeaders"  class="mtb20">
                    <el-form-item
                        label="header名称"
                        :key="'Container.'+index+'.lifecycle.postStart.httpGet.httpHeaders.'+index2+'.name'"
                        :prop="'Container.'+index+'.lifecycle.postStart.httpGet.httpHeaders.'+index2+'.name'"
                    >
                      <el-input v-model="data.name" placeholder="header名称"/>
                    </el-form-item>
                    <el-form-item
                        label="header值"
                        :key="'Container.'+index+'.lifecycle.postStart.httpGet.httpHeaders.'+index2+'.value'"
                        :prop="'Container.'+index+'.lifecycle.postStart.httpGet.httpHeaders.'+index2+'.value'"
                    >
                      <el-input v-model="data.value" placeholder="header值"/>
                    </el-form-item>
                    <el-form-item  >
                      <el-button size="small" type="danger" @click="removeHeader(index,index2,'lifecycle.postStart')" >删除</el-button>
                    </el-form-item>
                  </div>
                </div>
                <div v-if="item.lifecycle.postStart.type=='tcpSocket'">
                  <el-form-item
                      label="host"
                      :key="'Container.'+index+'.lifecycle.postStart.tcpSocket.host'"
                      :prop="'Container.'+index+'.lifecycle.postStart.tcpSocket.host'"
                  >
                    <el-input v-model="item.lifecycle.postStart.tcpSocket.host"  class="longWidth" placeholder="主机名称"/>
                  </el-form-item>
                  <el-form-item
                      label="主机端口"
                      :key="'Container.'+index+'.lifecycle.postStart.tcpSocket.port'"
                      :prop="'Container.'+index+'.lifecycle.postStart.tcpSocket.port'"
                  >
                    <el-input v-model="item.lifecycle.postStart.tcpSocket.port"  class="longWidth" placeholder="主机port"/>
                  </el-form-item>
                </div>
              </div>
              <el-divider></el-divider>
              <h3 >生命周期钩子preStop<el-button size="small" type="success"  class="mf20" @click="setResourcesPreStop(index)" >添加preStop</el-button></h3>
              <div v-if="item.lifecycle&&item.lifecycle.preStop" class="mtb20">
                <el-form-item label="设置探针类型"
                              :key="'Container.'+index+'.lifecycle.preStop.type'"
                              :prop="'Container.'+index+'.lifecycle.preStop.type'"
                >
                  <el-select  size="small" v-model="item.lifecycle.preStop.type"  class="select">
                    <el-option label="exec"  value="exec" />
                    <el-option label="httpGet"   value="httpGet"  />
                    <el-option label="tcpSocket"   value="tcpSocket"  />
                  </el-select>
                </el-form-item>
                <el-form-item  >
                  <el-button size="small" type="danger" @click="removeResources(index,'lifecycle','preStop')" >取消preStop</el-button>
                </el-form-item>
                <div v-if="item.lifecycle.preStop.type=='exec'">
                  <el-form-item
                      label="执行命令"
                      :key="'Container.'+index+'.lifecycle.preStop.exec.command'"
                      :prop="'Container.'+index+'.lifecycle.preStop.exec.command'"
                  >
                    <el-input v-model="item.lifecycle.preStop.exec.command"  class="longWidth" placeholder="执行命令"/>
                  </el-form-item>
                  <span  class="tip_info">多个值空格分割,并且空格使用双引号或者单引号包裹，例如 echo -c 'echo The app is running! && sleep 3600' </span>
                </div>
                <div v-if="item.lifecycle.preStop.type=='httpGet'">
                  <el-form-item
                      label="Host"
                      :key="'Container.'+index+'.lifecycle.preStop.httpGet.host'"
                      :prop="'Container.'+index+'.lifecycle.preStop.httpGet.host'"
                  >
                    <el-input v-model="item.lifecycle.preStop.httpGet.host"  class="longWidth" placeholder="host"/>
                  </el-form-item>
                  <el-form-item
                      label="端口"
                      :key="'Container.'+index+'.lifecycle.preStop.httpGet.port'"
                      :prop="'Container.'+index+'.lifecycle.preStop.httpGet.port'"
                  >
                    <el-input v-model="item.lifecycle.preStop.httpGet.port"  class="longWidth" placeholder="port"/>
                  </el-form-item>
                  <el-form-item
                      label="路径"
                      :key="'Container.'+index+'.lifecycle.preStop.httpGet.path'"
                      :prop="'Container.'+index+'.lifecycle.preStop.httpGet.path'"
                  >
                    <el-input v-model="item.lifecycle.preStop.httpGet.path"  class="longWidth" placeholder="路径"/>
                  </el-form-item>
                  <el-form-item label="协议类型"
                                :key="'Container.'+index+'.lifecycle.preStop.httpGet.scheme'"
                                :prop="'Container.'+index+'.lifecycle.preStop.httpGet.scheme'"
                  >
                    <el-select  size="small" v-model="item.lifecycle.preStop.httpGet.scheme"  class="select">
                      <el-option label="https"  value="https" />
                      <el-option label="http"   value="http"  />
                    </el-select>
                  </el-form-item>
                  <el-form-item  >
                    <el-button size="small" type="success" @click="addHttpHeaders(index,'lifecycle.preStop')" >添加header</el-button>
                  </el-form-item>
                  <div v-for="(data,index2) in item.lifecycle.preStop.httpGet.httpHeaders"  class="mtb20">
                    <el-form-item
                        label="header名称"
                        :key="'Container.'+index+'.lifecycle.preStop.httpGet.httpHeaders.'+index2+'.name'"
                        :prop="'Container.'+index+'.lifecycle.preStop.httpGet.httpHeaders.'+index2+'.name'"
                    >
                      <el-input v-model="data.name" placeholder="header名称"/>
                    </el-form-item>
                    <el-form-item
                        label="header值"
                        :key="'Container.'+index+'.lifecycle.preStop.httpGet.httpHeaders.'+index2+'.value'"
                        :prop="'Container.'+index+'.lifecycle.preStop.httpGet.httpHeaders.'+index2+'.value'"
                    >
                      <el-input v-model="data.value" placeholder="header值"/>
                    </el-form-item>
                    <el-form-item  >
                      <el-button size="small" type="danger" @click="removeHeader(index,index2,'lifecycle.preStop')" >删除</el-button>
                    </el-form-item>
                  </div>
                </div>
                <div v-if="item.lifecycle.preStop.type=='tcpSocket'">
                  <el-form-item
                      label="host"
                      :key="'Container.'+index+'.lifecycle.preStop.tcpSocket.host'"
                      :prop="'Container.'+index+'.lifecycle.preStop.tcpSocket.host'"
                  >
                    <el-input v-model="item.lifecycle.preStop.tcpSocket.host"  class="longWidth" placeholder="主机名称"/>
                  </el-form-item>
                  <el-form-item
                      label="主机端口"
                      :key="'Container.'+index+'.lifecycle.preStop.tcpSocket.port'"
                      :prop="'Container.'+index+'.lifecycle.preStop.tcpSocket.port'"
                  >
                    <el-input v-model="item.lifecycle.preStop.tcpSocket.port"  class="longWidth" placeholder="主机port"/>
                  </el-form-item>
                </div>

              </div>
          </div>
            <div class="redbar"></div>
          <el-divider></el-divider>
        </div>
      </el-form>
    </div>
  </el-card>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {requireRules,inArrayWithMsg} from "../../helper/rules.ts"
import {pvcAllByNs} from "../../api/token/pvc";
import {secretAllByNs} from "../../api/token/secret/secret";
import {configmapAllByNs} from "../../api/token/configmap/configmap";
import {parseData, UnParseData} from "../../helper/helper";
import md5 from 'js-md5';
import {isArray} from "element-plus/lib/utils/util";
export default defineComponent({
  name: 'Container',
  emits:["input"],
  props:["namespace","tip"],
  components:{
  },
  setup: (props,{emit}) => {
    let count=0;
    const formRef=ref(null)
    let portsRef=ref([])
    let state=reactive({
      show:false,
      advanced:[],
      tip:props.tip,
      secret:[],
      configMap:[],
      persistentVolumeClaim:[],
      form:{
        Container:[]
      },
    })
    fetchData(props.namespace)
    function removeContainers(index){
      state.form.Container.splice(index,1)
    }
    //添加容器
    function addContainers(){
      state.show=true
      //默认 hostPath 类型
      state.form.Container.push({
        "name":"",
        "image":"",
        "command":"",
        "args":"",
        "workingDir":"/",
        "ports":[],
        "envFrom":[],
        "env":[],
        "resources":{},
        "volumeMounts":[],
        "livenessProbe":{},
        "readinessProbe":{},
        "startupProbe":{},
        "lifecycle":{},
        terminationMessagePath:"",
        terminationMessagePolicy:{},
        imagePullPolicy:{},
        securityContext:{}
      })
    }
    function addPort(index){
      state.form.Container[index].ports.push({
        hostPort:"",
        name:"",
        hostIP:"",
        containerPort:"",
        protocol:"TCP"
      })
    }
    function removeResources(index,key,subkey){
      if(subkey){
        delete  state.form.Container[index][key][subkey]
        return
      }
      state.form.Container[index][key]={}
    }
    function setResources(index){
      if(!state.form.Container[index].resources.limits){
        state.form.Container[index].resources={
          limits:{
            cpu:"",
            memory:""
          },
          requests:{
            cpu:"",
            memory:""
          }
        }
      }

    }
    function setStartupProbe(index){
      if(!state.form.Container[index].startupProbe.type){

        state.form.Container[index].startupProbe={
          initialDelaySeconds:1,
          timeoutSeconds:10,
          periodSeconds:1,
          successThreshold:3,
          failureThreshold:3,
          terminationGracePeriodSeconds:60,
          exec:{
            command:"",
          },
          httpGet:{
            path:"",
            port:"",
            host:"",
            scheme:"http",
            httpHeaders:[
            ]
          },
          tcpSocket:{},
          type:"httpGet"
        }
      }
    }
    function setLivenessProbe(index){
      if(!state.form.Container[index].livenessProbe.type){

        state.form.Container[index].livenessProbe={
          initialDelaySeconds:1,
          timeoutSeconds:10,
          periodSeconds:1,
          successThreshold:3,
          failureThreshold:3,
          terminationGracePeriodSeconds:60,
          exec:{
            command:"",
          },
          httpGet:{
            path:"",
            port:"",
            host:"",
            scheme:"http",
            httpHeaders:[
            ]
          },
          tcpSocket:{},
          type:"httpGet"
        }
      }
    }
    function setReadinessProbe(index){
      if(!state.form.Container[index].readinessProbe.type){

        state.form.Container[index].readinessProbe={
          initialDelaySeconds:1,
          timeoutSeconds:10,
          periodSeconds:1,
          successThreshold:3,
          failureThreshold:3,
          terminationGracePeriodSeconds:60,
          exec:{
            command:"",
          },
          httpGet:{
            path:"",
            port:"",
            host:"",
            scheme:"http",
            httpHeaders:[
            ]
          },
          tcpSocket:{},
          type:"httpGet"
        }
      }
    }

    function addHttpHeaders(index,name) {
      let action=state.form.Container[index]
      if(name.indexOf(".")>-1){
        let arr=name.split(".")
        for (let i=0;i<arr.length;i++){
          action=action[arr[i]]
        }
      }else{
        action=action[name]
      }
      action.httpGet.httpHeaders.push( {name:"",value:""})
    }
    function addEnvFrom(index){
      state.form.Container[index].envFrom.push({"prefix":"","configMapRef":{"name":"","optional":false},"secretRef":{"name":"","optional":false}})
    }
    function addEnv(index){
      state.form.Container[index].env.push({
        "name":"",
        "value":"",
        valueFrom:{
          fieldRef:{
            apiVersion:"",
            FieldPath:""
          },
          resourceFieldRef:{
            containerName:"",
            resource:"",
            divisor:1
          },
          configMapKeyRef:{
            name:"",
            key:"",
            optional:false
          },
          secretKeyRef:{
            name:"",
            key:"",
            optional:false
          }
        }
      })
    }

    function getObject(data,index,index2,name){
      state.form.Container[index][name][index2]=data
    }
    async function fetchData(ns){
      // alert(ns)
      try {
        let pvc=await  pvcAllByNs(ns)
        let secret=await  secretAllByNs(ns)
        let configmap=await  configmapAllByNs(ns)
        //console.log(pvc.data.data)
        state.persistentVolumeClaim=pvc.data.data
        state.secret=secret.data.data
        state.configMap=configmap.data.data
        clearInput()
      }catch (e){
        console.log(e)
      }

    }
    function clearInput(){
      let arr=[]
      state.form.Container.forEach((item,index)=>{
        item.envFrom=[]
      })
      state.form.Container.forEach((item,index)=>{
        item.env=[]
      })

      // commit()
    }
    watch(()=>props.namespace,()=>{
      fetchData(props.namespace)
    })

    function setData(list) {
      let data=[]
      list.forEach((item,index)=>{
        let result={}
        result.name=item.name
        result.image=item.image

        result.command=UnParseData(item.command)
        result.args=UnParseData(item.args)
        result.workingDir=item.workingDir
        result.ports=item.ports
        result.envFrom=UnFilterEnvFrom(item.envFrom)
        result.env=UnFilterEnv(item.env)
        result.resources=UnFilterResources(item.resources)
        result.volumeMounts=UnFilterVolumeMounts(item.volumeMounts)
        result.livenessProbe=UnFilterProbe(item.livenessProbe,index)
        result.readinessProbe=UnFilterProbe(item.readinessProbe,index)
        result.startupProbe=UnFilterProbe(item.startupProbe,index)
        result.lifecycle={}
        if(item.lifecycle&&item.lifecycle.postStart){
          result.lifecycle.postStart=UnFilterHandle(item.lifecycle.postStart,index)
        }
        if(item.lifecycle&&item.lifecycle.preStop){
          result.lifecycle.preStop=UnFilterHandle(item.lifecycle.preStop,index)
        }


        data.push(result)
      })
      state.form.Container=data
      commit()
    }
    function UnFilterHandle(item,index){
      let result={
        exec:{
          command:"",
        },
        httpGet:{
          path:"",
          port:"",
          host:"",
          scheme:"http",
          httpHeaders:[
          ]
        },
        tcpSocket:{},
        type:"httpGet"
      }
      if(!item.exec&&!item.tcpSocket&&!item.httpGet){
        return  {}
      }
      if(item.exec){
        result.type="exec"
        result.exec.command=UnParseData(item.exec.command)
      }
      if(item.tcpSocket){
        result.type="tcpSocket"
        result.tcpSocket=item.tcpSocket
      }
      if(item.httpGet){
        result.type="httpGet"
        result.httpGet=item.httpGet
      }
      return result
    }
    function UnFilterProbe(item,index) {
        let result={
          initialDelaySeconds:1,
          timeoutSeconds:10,
          periodSeconds:1,
          successThreshold:3,
          failureThreshold:3,
          terminationGracePeriodSeconds:60,
          exec:{
            command:"",
          },
          httpGet:{
            path:"",
            port:"",
            host:"",
            scheme:"http",
            httpHeaders:[
            ]
          },
          tcpSocket:{},
          type:"httpGet"
        }
      let attr=["initialDelaySeconds","timeoutSeconds","periodSeconds","successThreshold","failureThreshold","terminationGracePeriodSeconds"]
      for (let i in attr){
        if(item[attr[i]]){
          result[attr[i]]=item[attr[i]]
        }
      }
      if(!item.exec&&!item.tcpSocket&&!item.httpGet){
        return  {}
      }
      if(item.exec){
          result.type="exec"
          result.exec.command=UnParseData(item.exec.command)
      }
      if(item.tcpSocket){
        result.type="tcpSocket"
        result.tcpSocket=item.tcpSocket
      }
      if(item.httpGet){
        result.type="httpGet"
        result.httpGet=item.httpGet
      }
      return result
    }
    function filterEnv(data) {
      let result=[];
      data.forEach(function (item) {
        let temp={name:item.name}
        if( item.value){
          temp.value=item.value
          result.push(temp)
          return
        }
        temp.valueFrom={}
        if(item.valueFrom.fieldRef&&item.valueFrom.fieldRef.apiVersion||item.valueFrom.fieldRef.FieldPath){
          temp.valueFrom.fieldRef=item.valueFrom.fieldRef
        }
        if(item.valueFrom.resourceFieldRef&&item.valueFrom.resourceFieldRef.containerName||item.valueFrom.resourceFieldRef.resource){
          temp.valueFrom.resourceFieldRef=item.valueFrom.resourceFieldRef
        }
        if(item.valueFrom.configMapKeyRef&&item.valueFrom.configMapKeyRef.name||item.valueFrom.configMapKeyRef.key){
          temp.valueFrom.configMapKeyRef=item.valueFrom.configMapKeyRef
        }
        if(item.valueFrom.secretKeyRef&&item.valueFrom.secretKeyRef.name||item.valueFrom.secretKeyRef.key){
          temp.valueFrom.secretKeyRef=item.valueFrom.secretKeyRef
        }
        result.push(temp)
      })
      return result
    }
    function filterEnvFrom(data) {
      let result=[];
      data.forEach(function (item) {
        let temp={prefix:item.prefix}
        if(item.configMapRef&&item.configMapRef.name){
          temp.configMapRef=item.configMapRef
        }
        if(item.secretRef&&item.secretRef.name){
          temp.secretRef=item.secretRef
        }
        result.push(temp)
      })
      return result
    }
    function UnFilterEnvFrom(data){
      let result=[];
      data.forEach(function (item) {
        let temp={prefix:item.prefix,configMapRef:{"name":"","optional":false},secretRef:{"name":"","optional":false}}
        if(item.configMapRef){
          temp.configMapRef = item.configMapRef
        }
        if(item.secretRef){
          temp.secretRef=item.secretRef
        }
        result.push(temp)
      })
      return result
    }
    function UnFilterEnv(data){
      let result=[];
      data.forEach(function (item) {
        let temp={valueFrom:{
            fieldRef:{
              apiVersion:"",
              FieldPath:""
            },
            resourceFieldRef:{
              containerName:"",
              resource:"",
              divisor:1
            },
            configMapKeyRef:{
              name:"",
              key:"",
              optional:false
            },
            secretKeyRef:{
              name:"",
              key:"",
              optional:false
            }
            },name:item.name,value:""}

        if(item.valueFrom&&item.valueFrom.fieldRef){
          temp.valueFrom.fieldRef=item.valueFrom.fieldRef
        }
        if(item.value){
          temp.value=item.value
        }
        if(item.valueFrom&&item.valueFrom.resourceFieldRef){
          temp.valueFrom.resourceFieldRef=item.valueFrom.resourceFieldRef
        }
        if(item.valueFrom&&item.valueFrom.configMapKeyRef){
          temp.valueFrom.configMapKeyRef=item.valueFrom.configMapKeyRef
        }
        if(item.valueFrom&&item.valueFrom.secretKeyRef){
          temp.valueFrom.secretKeyRef=item.valueFrom.secretKeyRef
        }

        result.push(temp)
      })
      return result
    }
    function commit() {
      let data=[]
      state.form.Container.forEach(item=>{
        let result={}
        result.name=item.name
        result.image=item.image

        result.command=parseData(item.command)


        result.args=parseData(item.args)

        result.workingDir=item.workingDir
        result.ports=item.ports
        result.envFrom=filterEnvFrom(item.envFrom)
        result.env=filterEnv(item.env)
        result.resources=filterResources(item.resources)
        result.volumeMounts=filterVolumeMounts(item.volumeMounts)
        result.livenessProbe=filterProbe(item.livenessProbe)
        result.readinessProbe=filterProbe(item.readinessProbe)
        result.startupProbe=filterProbe(item.startupProbe)
        result.lifecycle={}
        if(item.lifecycle&&item.lifecycle.postStart){
          result.lifecycle.postStart=filterHandle(item.lifecycle.postStart)
        }
        if(item.lifecycle&&item.lifecycle.preStop){
          result.lifecycle.preStop=filterHandle(item.lifecycle.preStop)
        }

        data.push(result)
      })

      emit("input",data)
    }
    function filterHandle(item){
      let result={}
      if(!item||!item.type){
        return result
      }
      if(item.type=="exec"){
        result.exec={}
        result.exec.command=parseData(item.exec.command)
        console.log(count++)
      }
      if(item.type=="httpGet"){
        result.httpGet=item.httpGet
      }
      if(item.type=="tcpSocket"){
        result.tcpSocket=item.tcpSocket
      }
      return result
    }
    function UnFilterResources(item) {
      let result={
        limits:{
          cpu:"",
          memory:""
        },
        requests:{
          cpu:"",
          memory:""
        }
      }
      let count=1;
      if(item.limits&&item.limits.cpu){
        result.limits.cpu=item.limits.cpu
        count--
      }
      if(item.limits&&item.limits.memory){
        result.limits.memory=item.limits.memory
        count--
      }
      if(item.requests&&item.requests.cpu){
        result.requests.cpu=item.requests.cpu
        count--
      }
      if(item.requests&&item.requests.memory){
        result.requests.memory=item.requests.memory
        count--
      }
      if(count==1){
        result={}
      }
      return result

    }
    function filterResources(item) {
      let result={}
      if(item.limits&&item.limits.cpu){
        if(!result.limits){
          result.limits={}
        }
        result.limits.cpu=item.limits.cpu
      }
      if(item.limits&&item.limits.memory){
        if(!result.limits){
          result.limits={}
        }
        result.limits.memory=item.limits.memory
      }
      if(item.requests&&item.requests.cpu){
        if(!result.requests){
          result.requests={}
        }
        result.requests.cpu=item.requests.cpu
      }
      if(item.requests&&item.requests.memory){
        if(!result.requests){
          result.requests={}
        }
        result.requests.memory=item.requests.memory
      }
      return result
    }
    function filterProbe(item) {
      let result={}
      if(!item||!item.type){
        return result
      }
      let attr=["initialDelaySeconds","timeoutSeconds","periodSeconds","successThreshold","failureThreshold","terminationGracePeriodSeconds"]
      for (let i in attr){
        if(item[attr[i]]){
          result[attr[i]]=item[attr[i]]
        }
      }
      if(item.type=="exec"){
        result.exec={}
        result.exec.command=parseData(item.exec.command)
        console.log(count++)
      }
      if(item.type=="httpGet"){
        result.httpGet=item.httpGet
      }
      if(item.type=="tcpSocket"){
        result.tcpSocket=item.tcpSocket
      }
      return result
    }
    watch(()=>state.form,()=>{
      commit()
    },{deep:true,flush:"post"})
    function removeItem(index,index2,name) {

      state.form.Container[index][name].splice(index2,1)
    }
    function removeHeader(index,index2,name) {
      let action=state.form.Container[index]
      if(name.indexOf(".")>-1){
        let arr=name.split(".")
        for (let i=0;i<arr.length;i++){
          action=action[arr[i]]
        }
      }else{
        action=action[name]
      }
      action.httpGet.httpHeaders.splice(index2,1)
    }
    function UnFilterVolumeMounts(list) {
      let arr=[]
      list.forEach((item)=>{
        let result={
          name:"",
          readOnly:false,
          mountPath:"",
          subPath:"",
          subPathExpr:"",
          mountPropagation:"None"
        }
        let attr=["name","readOnly","mountPath","subPath","subPathExpr","mountPropagation"]
        for (let i in attr){
          if(item[attr[i]]){
            result[attr[i]]=item[attr[i]]
          }
        }
        arr.push(result)
      })
       return arr
    }
    function addVolumeMounts(index) {
      state.form.Container[index].volumeMounts.push({
        name:"",
        readOnly:false,
        mountPath:"",
        subPath:"",
        subPathExpr:"",
        mountPropagation:"None"
      })
    }

    function filterVolumeMounts(list) {
      let arr=[]
      list.forEach((item)=>{
        let result={}
        let must=["name","mountPath"]
        let option=["name","readOnly","mountPropagation"]
        for (let i in must){
            result[must[i]]=item[must[i]]
        }
        for (let i in option){
          if(item[option[i]]){
            result[option[i]]=item[option[i]]
          }
        }
        if(item['subPath']){
          result['subPath']=item['subPath']
        }
        if(!item['subPath']&&item['subPathExpr']){
          result['subPathExpr']=item['subPathExpr']
        }
        arr.push(result)
      })
      return arr
    }
    function setResourcesPostStart(index) {

          state.form.Container[index].lifecycle.postStart={
            exec:{
              command:"",
            },
            httpGet:{
              path:"",
              port:"",
              host:"",
              scheme:"http",
              httpHeaders:[
              ]
            },
            tcpSocket:{},
            type:"httpGet"
        }
    }
    function setResourcesPreStop(index){
      state.form.Container[index].lifecycle.preStop={
        exec:{
          command:"",
        },
        httpGet:{
          path:"",
          port:"",
          host:"",
          scheme:"http",
          httpHeaders:[
          ]
        },
        tcpSocket:{},
        type:"httpGet"
      }
    }
    function setAdvanced(i){
      let temp=state.advanced[i]
     for (let k=0;k<state.advanced.length;k++){
       state.advanced[k]=false
     }
     state.advanced[i]=!temp
    }
    return {...toRefs(state),formRef,setAdvanced,setResourcesPreStop,setResourcesPostStart,setStartupProbe,setReadinessProbe,addContainers,requireRules,addPort,getObject,addEnvFrom,addEnv,setData,parseData,removeContainers,removeItem,setResources,setLivenessProbe,removeResources,addVolumeMounts,addHttpHeaders,removeHeader}
  }
})
</script>

<style scoped>

</style>
