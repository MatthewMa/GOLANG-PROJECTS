/**
 * Created by Sihua on 2017/2/10.
 */
/// <reference path="../angular.js" />

angular.module("sportsStore")
    .constant("dataUrl", "http://localhost:2403/products")
    .constant("orderUrl","http://localhost:2403/order")
    .controller("sportsStoreCtrl", function ($scope, $http, $location,dataUrl,orderUrl,cart) {

        $scope.data = {};
        $http.get(dataUrl)
            .then(function (result) {
                console.info(result.data)
                $scope.data.products = result.data;
            }).catch(function (result) {
                console.info(result);
                $scope.data.error=result.data;
            });
        $scope.sendOrder = function (shippingDetails) {
            var order = angular.copy(shippingDetails);
            order.products = cart.getProducts();
            $http.post(orderUrl, order)
                .then(function successCallback (response) {
                    $scope.data.orderId = response.data.id;
                    cart.getProducts().length = 0;
                },function errorCallback(response) {
                    $scope.data.orderError = response.status;
                }).finally(function () {
                    $location.path('/complete');
                });
        }
    });
