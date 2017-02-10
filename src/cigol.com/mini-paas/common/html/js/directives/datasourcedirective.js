define(['app','jquery'],function(app,$){
        //表单 input-下拉树选择组件
        app.register.directive("comborTree", function($global_var,$q,$http) {
            return {
                restrict: 'E',
                transclude:true,
                templateUrl:'templates/app/pluginTemplates/combor-tree.html',
                link: function(scope, element, attr) {
                    //--------------------scope初始化部分-----------------------------------------------
                    if(attr.messages){
                        scope.messageList = JSON.parse(attr.messages)
                    }
                    scope.rows=[];
                    scope.openTreeClick=false;
                    scope.initTree=function(){
                        $.fn.zTree.init($('#combortreeId'), comborsetting, []);
                        scope.zztreeObj = $.fn.zTree.getZTreeObj("combortreeId");
                        scope.zztreeObj.addNodes(null,scope.rows);
                        scope.zztreeObj.expandAll(true);
                    };

                    //新建
                    scope.getTreeData = function(){
                        var deferred=$q.defer();
                        var deferred=$q.defer();
                        $http({
                            method: 'GET',
                            url:$global_var.base_url +attr.url+"?id="+scope.currentTableDataSourceId,
                        }).success(function(result) {
                            deferred.resolve(result);
                        }).error(function(err){
                            deferred.reject(false);
                            console.log(err);
                        });
                        return deferred.promise;
                    };

                    //---------------------tree操作部分--------------------------------------------------
                    var comborsetting = {
                        view:{
                            showIcon:false
                        },
                        check:{
                            enable:true,
                            chkboxType: { "Y": "", "N": "" }
                        },
                        callback:{
                            onCheck:onCheckTree
                        }
                    };
                    function onCheckTree(e, treeId, treeNode) {
                        var showTableName = "";
                        var selectNodes = scope.zztreeObj.getCheckedNodes(true);
                        if(selectNodes && selectNodes.length>0){
                            angular.forEach(selectNodes, function(dataNode,index,array){
                                showTableName += (dataNode.id + ',')
                            });
                            showTableName = showTableName.slice(0, showTableName.length - 1)
                        }
                        scope.selectTableName=showTableName;
                        scope.$apply();  //外部dom改变scope，手动调用$apply来刷新加载scope
                    }

                    //----------------------dom操作部分----------------------------------------------------
                    scope.showSearch=function(obj){
                        scope.openTreeClick=true;
                        scope.getTreeData().then(function (result) {
                            scope.rows = angular.fromJson(result.content);
                            scope.initTree();
                        });

                    }
                    scope.hideSearch=function(){
                        if(scope.zztreeObj)
                            scope.zztreeObj.destroy();
                        $('#dataSourceChoose').modal('hide');
                        scope.openTreeClick=false;
                    }
                }
            }
        });


    })