/**
 * Created by Sihua on 2017/2/17.
 */
angular.module("sportsStoreAdmin")
    .constant("authUrl", "http://localhost:2403/users/login")
    .controller("authCtrl", function($scope, $http, $location, authUrl) {

        $scope.authenticate = function (user, pass) {
            $http.post(authUrl, {
                username: user,
                password: pass
            }, {
                withCredentials: true
            }).success(function (data) {
                $scope.goPath = function(){
                    $location.path('/main');
                }
            }).error(function (error) {
                $scope.authenticationError = error;
            });
        }
    })
    .controller("mainCtrl", function($scope) {

        $scope.screens = ["Products", "Orders"];
        $scope.current = $scope.screens[0];

        $scope.setScreen = function (index) {
            $scope.current = $scope.screens[index];
        };

        $scope.getScreen = function () {
            return $scope.current == "Products"
                ? "/static/js/views/adminProducts.html" : "/static/js/views/adminOrders.html";
        };
    });
