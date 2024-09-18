#!/usr/bin/env bash

function deleteCode() {
  fromLine=$1
  shouldDelLines=$2
  fineName=$3

  d=$(expr $fistIndex + $shouldDelLines)
  insertExpr=$(echo "${fistIndex},${d} d")
  sed -i "$insertExpr" $fineName
}

function getdir() {
  for file in $1/*; do
    if test -f "$file"; then
      if [ "${file##*.}"x = "go"x ]; then
        #        if [ $(basename $file) == 'router.go' ]; then
#        arr=(${arr[*]} ${projectPath#$file})
        arr=(${arr[*]} $file)
        #        fi
      fi
    else
      getdir $file
    fi
  done
}

function router() {
  g=$1
  if [ $(basename $g) == 'router.go' ]; then
    sed -i '4i\"github.com/pengcainiao2/zero/rest"' "$g"
    sed -i 's/gin.Default()/rest.NewGinServer()/g' "$g"
    # 删除graceFullShutdown中对health server的引用
    sed -i 's/setupHealthZ(),//' "$g"
    sed -i 's/healthSrv,//' "$g"
    #获取healthSrv.Shutdown位置并删除
    fistIndex=$(grep -n 'healthSrv.Shutdown' "$g" | cut -d ":" -f 1)
    deleteCode $fistIndex 2 "$g"
    fistIndex=$(grep -n 'func setupHealthZ()' "$g" | cut -d ":" -f 1)
    deleteCode $fistIndex 16 "$g"
  fi
}

function insertImports() {
  g=$1
  cmd=$2
  fistIndex=$(grep -n 'import (' "$g" | cut -d ":" -f 1)

#  echo "当前处理："$g $fistIndex
  if [ $fistIndex -gt 0 ] && [ $fistIndex != '' ]; then
       d=$(expr $fistIndex + 1)
          insertExpr=$(echo "${d}i\\"'"'$cmd'"')
          sed -i $insertExpr "$g"
  fi
}

function replaceGoImport() {
  g=$1
  if [ -f "$g" ]; then


  insertImports "$g" "github.com/pengcainiao2/zero/core/logx"
  insertImports "$g" "github.com/pengcainiao2/zero/core/queue/nsqueue"
  insertImports "$g" "github.com/pengcainiao2/zero/core/snowflake"
  insertImports "$g" "github.com/pengcainiao2/zero/core/sensorsx"
  insertImports "$g" "github.com/pengcainiao2/zero/core/sensorsx/items"
  insertImports "$g" "github.com/pengcainiao2/zero/core/discov"
  insertImports "$g" "gitlab.flyele.vip/server-side/go-zero/core/timex"
  insertImports "$g" "gitlab.flyele.vip/server-side/go-zero/core/conf"

  sed -i 's/gitlab.flyele.vip\/server-side\/tools\/httprouter/gitlab.flyele.vip\/server-side\/go-zero\/v2\/rest\/httprouter/g' "$g"
  sed -i 's/gitlab.flyele.vip\/server-side\/tools\/syncer/gitlab.flyele.vip\/server-side\/go-zero\/v2\/tools\/syncer/g' "$g"
  sed -i 's/gitlab.flyele.vip\/server-side\/tools\/webhook/gitlab.flyele.vip\/server-side\/go-zero\/v2\/rest\/httpx/g' "$g"

  sed -i 's/gitlab.flyele.vip\/server-side\/tools\/sensorsdata\/events/gitlab.flyele.vip\/server-side\/go-zero\/v2\/tools\/sensors\/events/g' "$g"
  sed -i 's/gitlab.flyele.vip\/server-side\/tools\/sensorsdata/gitlab.flyele.vip\/server-side\/go-zero\/v2\/tools\/sensors/g' "$g"

  sed -i 's/github.com\/go-redis\/redis/github.com\/go-redis\/redis\/v8/g' "$g"
  sed -i 's/gitlab.flyele.vip\/server-side\/tools\/env/gitlab.flyele.vip\/server-side\/go-zero\/v2\/core\/env/g' "$g"

  sed -i 's/gitlab.flyele.vip\/server-side\/tools\/cronjob/gitlab.flyele.vip\/server-side\/go-zero\/v2\/rest\/cronjob/g' "$g"
  fi
}

function replaceCode() {
  g=$1 #
  if [ -f "$g" ]; then
  sed -i 's/syncer.RandomExpireSeconds/timex.RandomExpireSeconds/g' "$g"

  sed -i 's/zerolog.ErrorLogger(/logx.NewTraceLogger(context.Background()).Err(/g' "$g"
  sed -i 's/zerolog.Default()/logx.NewTraceLogger(context.Background())/g' "$g"
  sed -i 's/tools.GenerateID()/sonyflake.GenerateID()/g' "$g"
  sed -i 's/tools.GenerateInt64ID()/sonyflake.GenerateInt64ID()/g' "$g"
  # 更新redis,mysql
  sed -i 's/syncer.Redis()[A-Za-z\.]\{1,\}(\{1\}/&context.Background(),/' "$g"
  sed -i 's/redisClient[A-Za-z\.]\{1,\}(\{1\}/&context.Background(),/' "$g"
  sed -i 's/rdsClient[A-Za-z\.]\{1,\}(\{1\}/&context.Background(),/' "$g"
  sed -i 's/pipeline[A-Za-z\.]\{1,\}(\{1\}/&context.Background(),/' "$g"
  sed -i 's/syncer.MySQL()[A-Za-z\.]\{1,\}(\{1\}/&context.Background(),/' "$g"
  sed -i 's/webhook.ReportErrors/httpx.ReportErrors/g' "$g"
  sed -i 's/syncer.ElasticSearch()/search.ElasticSearch()/g' "$g"
  sed -i 's/syncer.WithMySQLMaxIdleConns/conf.WithMySQLMaxIdleConns/g' "$g"
  sed -i 's/syncer.WithMySQLMaxOpenConns/conf.WithMySQLMaxOpenConns/g' "$g"
  sed -i 's/syncer.WithMaxNSQGoroutineCount/conf.WithMaxNSQGoroutineCount/g' "$g"
  sed -i 's/syncer.WithMySQLMaxIdleDuration/conf.WithMySQLMaxIdleDuration/g' "$g"
  sed -i 's/syncer.WithNsqEndpoint/conf.WithNsqEndpoint/g' "$g"
  sed -i 's/syncer.WithEtcdEndpoint/conf.WithEtcdEndpoint/g' "$g"
  sed -i 's/syncer.WithMaxConcurrency/conf.WithMaxConcurrency/g' "$g"
  sed -i 's/syncer.WithRedisPoolSize/conf.WithRedisPoolSize/g' "$g"
  sed -i 's/syncer.WithRedisMinIdleConns/conf.WithRedisMinIdleConns/g' "$g"

# 更新es
  sed -i 's/syncer.IndexInteractions/search.IndexInteractions/g' "$g"
  sed -i 's/syncer.IndexConversation/search.IndexConversation/g' "$g"
  sed -i 's/syncer.IndexFlyeleTask/search.IndexFlyeleTask/g' "$g"
  sed -i 's/syncer.IndexAttachmentV1/search.IndexAttachmentV1/g' "$g"
  sed -i 's/syncer.IndexTimeRemind/search.IndexTimeRemind/g' "$g"
  sed -i 's/syncer.IndexDraft/search.IndexDraft/g' "$g"
  sed -i 's/syncer.IndexRecord/search.IndexRecord/g' "$g"
  sed -i 's/syncer.IndexTaskDispatch/search.IndexTaskDispatch/g' "$g"
  sed -i 's/syncer.IndexTask/search.IndexTask/g' "$g"

  # 更新NSQ队列
  sed -i 's/syncer.NsqTraceID/nsqueue.NsqTraceID/g' "$g"
  sed -i 's/syncer.Topic/nsqueue.Topic/g' "$g"
  sed -i 's/syncer.MqMessageConsumer/nsqueue.MqMessageConsumer/g' "$g"
  sed -i 's/syncer.NsqDataProtocol/nsqueue.NsqDataProtocol/g' "$g"
  sed -i 's/syncer.UnmarshalNsqMessage/nsqueue.UnmarshalNsqMessage/g' "$g"

  sed -i 's/syncer.NsqPushGatewayMulticastDataTopic/syncer.QueuePushGatewayMulticast().Priority(nsqueue.OrderedNormal)/g' "$g"
  sed -i 's/syncer.NsqTimedTaskDataTopic/syncer.QueueTimedTask().Priority(nsqueue.OrderedNormal)/g' "$g"
  sed -i 's/syncer.NsqDataSyncerSvcTopic/syncer.QueueDataSyncer().Priority(nsqueue.OrderedNormal)/g' "$g"

  sed -i 's/syncer.NsqPushGatewayUnicastDataTopic/syncer.QueuePushGatewayUnicast().Priority(nsqueue.NoOrderNormal)/g' "$g"
  sed -i 's/syncer.NsqTimedTaskNoneOrderDataTopic/syncer.QueueTimedTask().Priority(nsqueue.NoOrderNormal)/g' "$g"
  sed -i 's/syncer.NsqUserInteractionNoneOrderSvcTopic/syncer.QueueUserInteraction().Priority(nsqueue.NoOrderNormal)/g' "$g"
  sed -i 's/syncer.NsqCloudiskNoneOrderSvcTopic/syncer.QueueCloudDisk().Priority(nsqueue.NoOrderNormal)/g' "$g"
  sed -i 's/syncer.NsqFlyeleSvcLowPriorityTopic/syncer.QueueFlyele().Priority(nsqueue.NoOrderLow)/g' "$g"
  sed -i 's/syncer.NsqFlyeleSvcTopic/syncer.QueueFlyele().Priority(nsqueue.NoOrderNormal)/g' "$g"
  sed -i 's/syncer.NsqRecordNoneOrderSvcTopic/syncer.QueueRecord().Priority(nsqueue.NoOrderNormal)/g' "$g"
  sed -i 's/syncer.NsqUserCenterSvcTopic/syncer.QueueRecord().Priority(nsqueue.NoOrderNormal)/g' "$g"

  # senseors
  sed -i 's/sensorsdata.Client()/sensors.SensorsClient()/g' "$g"
  sed -i 's/sensorsdata.BaseEvent/sensorsx.BaseEvent/g' "$g"
  sed -i 's/sensorsdata.SensorTask/sensorsx.SensorTask/g' "$g"
  sed -i 's/sensorsdata.Normal/sensorsx.Normal/g' "$g"
  sed -i 's/sensorsdata.SignUp/sensorsx.SignUp/g' "$g"
  sed -i 's/sensorsdata.UpdateUserProfile/sensorsx.UpdateUserProfile/g' "$g"
  sed -i 's/sensorsdata.ItemSet/sensorsx.ItemSet/g' "$g"

  #etcd
  sed -i 's/syncer.EtcdChangeType/discov.EtcdChangeType/g' "$g"
  sed -i 's/syncer.GetResponse/discov.GetResponse/g' "$g"
  sed -i 's/syncer.EtcdClient/*discov.EtcdClient/g' "$g"

# httpx
sed -i 's/webhook.CustomerFeedback/*httpx.CustomerFeedback/g' "$g"
sed -i 's/webhook.TxtMessage/*httpx.TxtMessage/g' "$g"
sed -i 's/webhook.MarkdownMessage/*httpx.MarkdownMessage/g' "$g"
sed -i 's/tools.SendHTTP/*httpx.SendHTTP/g' "$g"
  fi
}

projectPath=/Users/rudy/Desktop/work/coding/golang/src/gitlab.flyele.vip/flyele
getdir $projectPath

cd $projectPath
go get -u github.com/pengcainiao2/zero

for g in ${arr[*]}; do
  router ""$g""
  replaceGoImport "$g"
  replaceCode ""$g""
  #    go fmt "$g"
  goimports -w "$g"
done

go mod tidy
