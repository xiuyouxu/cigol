/**
 * Created by wei.fan on 2016/11/21.
 */
define(['app'],function(app){
    app.register.service('$appIndexService', ['$http','$q','$global_var',function($http,$q,$global_var){

        //根据Id查询用户信息
        this.queryUserById = function(data){
            var deferred=$q.defer();
            $http({
                method: 'GET',
                url: $global_var.base_url + "base/user/"+data
            }).success(function(result) {
                deferred.resolve(result);
            }).error(function(err){
                deferred.reject(false);
                console.log(err);
            });
            return deferred.promise;
        };

        //编辑
        this.edit = function(page){
            var deferred=$q.defer();
            $http({
                method: 'POST',
                url:$global_var.base_url + "base/user/updating" ,
                data:page
            }).success(function(result) {
                deferred.resolve(result);
            }).error(function(err){
                deferred.reject(false);
                console.log(err);
            });
            return deferred.promise;
        };


        //用户名手机号去重
        this.nameOrPhoneRepeat = function(data){
            var deferred=$q.defer();
            $http({
                method: 'POST',
                url: $global_var.base_url + "base/user/byCondition",
                data:data
            }).success(function(result) {
                deferred.resolve(result);
            }).error(function(err){
                deferred.reject(false);
                console.log(err);
            });
            return deferred.promise;
        };


    }]);
});
