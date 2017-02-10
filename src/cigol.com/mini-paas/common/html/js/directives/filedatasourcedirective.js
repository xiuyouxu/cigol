define(['app','jquery'],function(app,$){
        //表单 input-下拉树选择组件
        app.register.directive("fileComborTree", function($global_var,$q,$http) {
            return {
                restrict: 'E',
                transclude:true,
                templateUrl:'templates/app/pluginTemplates/file-combor-tree.html',
                link: function(scope, element, attr) {
                    //--------------------scope初始化部分-----------------------------------------------
                    scope.dirRows=[];
                    scope.openDirTreeClick=false;
                    scope.initDirTree=function(){

                        var rootNode = {"id":"0", "name":"目录", "open":true, "isParent":true};
                        $.fn.zTree.init($('#filecombortreeId'), dirTreeSetting, rootNode);
                        scope.dirztreeObj = $.fn.zTree.getZTreeObj("filecombortreeId");
                        // var node = zTreeObj.getNodeByParam("id", 0, null);
                        // zTreeObj.reAsyncChildNodes(node, "refresh");
                    };

                    //---------------------tree操作部分--------------------------------------------------
                    var dirTreeSetting = {
                        view:{
                            showIcon:false
                        },
                        async:{
                            enable: true,
                            url:$global_var.base_url+attr.url,
                            type: "get",
                            dataType:"json",
                            autoParam:["id=pId"]
                        },
                        check:{
                            enable:true,
                            chkboxType: { "Y": "", "N": "" }
                        },
                        callback:{
                            onCheck:onCheckDirTree
                        }
                    };
                    function idFilter(treeId, parentNode, childNodes) {
                        angular.forEach(childNodes, function(dataNode,index,array){
                            dataNode.id = dataNode.id.slice(1);
                        });
                    }
                    function onCheckDirTree(e, treeId, treeNode) {
                        var showTableName = "";
                        var showTableType = "";
                        var selectNodes = scope.dirztreeObj.getCheckedNodes(true);
                        if(selectNodes && selectNodes.length>0){
                            scope.currentCheckType = "1";
                            angular.forEach(selectNodes, function(dataNode,index,array){
                                if("0"===dataNode.id){
                                }else{
                                    if(scope.currentDataSourceType==="23"){
                                        showTableName += (dataNode.id + '/,');
                                    }else{
                                        showTableName += (dataNode.id + ',');
                                    }
                                    showTableType += (dataNode.isParent + ',');
                                }

                            });
                            showTableName = showTableName.slice(0, showTableName.length - 1);
                            showTableType = showTableType.slice(0, showTableType.length - 1);
                        }else{
                            scope.currentCheckType = "2";
                        }
                        scope.selectDirName = showTableName;
                        scope.selectDirType = showTableType;
                        scope.$apply();  //外部dom改变scope，手动调用$apply来刷新加载scope
                    }

                    //----------------------dom操作部分----------------------------------------------------
                    scope.showDirSearch=function(obj){
                        scope.openDirTreeClick=true;
                        dirTreeSetting.async.otherParam = { "id":scope.currentDirDataSourceId};
                        //？？？延迟加载初始化树
                        setTimeout(function(){
                            scope.initDirTree();
                        },300);

                    }
                    scope.hideDirSearch=function(){
                        if(scope.dirztreeObj)
                            scope.dirztreeObj.destroy();
                        $('#fileDirChoose').modal('hide');
                        scope.openDirTreeClick=false;
                    }
                }
            }
        });


    })