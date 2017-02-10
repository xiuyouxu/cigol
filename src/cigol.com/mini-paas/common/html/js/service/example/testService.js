define(['app'],function(app){
    app.register.service('$testService', ['$http','$q','$global_var',function($http,$q,$global_var){
        //查询树节点数据
        this.test = function () {
            // var deferred=$q.defer();
            // $http({
            //     method: 'GET',
            //     url: $global_var.base_url + "resource/catalogRestServer/tree?parentId=0&type=0&containLeaf=false"
            // }).success(function(result) {
            //     deferred.resolve(result);
            // }).error(function(err){
            //     deferred.reject(false);
            //     console.log(err);
            // });
            // return deferred.promise;
            return {
                'a': 'good',
                'b': 'you',
                'c': 'are',
                'd': 'welcome!'
            }
        }
        this.getHosts = function () {
            var deferred=$q.defer();
            $http({
                method: 'GET',
                url: $global_var.base_url + "hosts"
            }).success(function(result) {
                deferred.resolve(result);
            }).error(function(err){
                deferred.reject(false);
                console.log(err);
            });
            return deferred.promise;
        }
    }]);
});


