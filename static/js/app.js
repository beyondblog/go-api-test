var routerApp = angular.module('httpTest', ['ui.router']);


routerApp.factory('editApi', function() {
    var service = {};
    var requests;
    var editRequest;
    var hostName;

    service.setEditRequest = function(request) {
        this.editRequest = request;
    };

    service.getEditRequest = function() {
        return this.editRequest;
    };

    service.setRequests = function(requests) {
        this.requests = requests;
    };

    service.getRequests = function() {
        return this.requests;
    };

    service.setHostName = function(hostName) {
        this.hostName = hostName;
    };

    service.getHostName = function() {
        return this.hostName;
    };

    return service;
});

routerApp.config(function($stateProvider, $urlRouterProvider, $locationProvider) {

    $locationProvider.html5Mode({
        enabled: true,
        requireBase: false
    });

    $urlRouterProvider.otherwise('/');

    $stateProvider
        .state('add', {
            url: '/',
            templateUrl: '/views/add.html',
            controller: 'add'
        })
        .state('list', {
            url: '/list',
            templateUrl: '/views/list.html',
            controller: 'list'
        })
        .state('about', {
            url: '/about',
            templateUrl: '/views/about.html'
        }).state('edit', {
            url: '/edit',
            templateUrl: '/views/edit.html',
            controller: 'edit'
        }).state('editRequest', {
            url: '/edit/:id',
            templateUrl: '/views/editRequest.html',
            controller: 'editRequest'
        });

});

routerApp.controller('add', function($scope, $http, $location) {
    $scope.host = '';
    $scope.desc = '';
    $scope.message = '';
    $scope.method = 0;
    $scope.params = [{}];

    $scope.add = function() {
        $http.post('/api/add', {
            host: $scope.host,
            desc: $scope.desc,
            method: parseInt($scope.method),
            param: $scope.params,
        }).success(function(data) {
            $scope.message = data.message;
            if (data.Code == 200) {
                $scope.host = '';
                $scope.desc = '';
                $scope.message = '';
                $scope.method = 0;
                $scope.params = [{}];
            }
        }).error(function() {
            $scope.message = 'server error : (';
        });
    };

    $scope.keyvalueClick = function() {
        var last = $scope.params[$scope.params.length - 1];
        if (last.key != null || last.value != null) {
            $scope.params.push({})
        }
    };

    $scope.delParam = function() {
        var index = this.$index;
        if (~index) $scope.params.splice(index, 1);
    };

}).controller('list', function($scope, $http, $location, editApi) {
    $scope.requests = {};
    $scope.init = function() {
        $http.get('/api/list').success(function(data) {
            if (data.Code == 200) {
                $scope.requests = data.Data;
            }
        }).error(function() {
            $scope.message = 'server error : (';
        });
    };

    $scope.edit = function() {
        editApi.setHostName(this.$$watchers[1].last);
        editApi.setRequests($scope.requests[this.$index].Requests);
        $location.path("/edit");
    };

    $scope.init();

}).controller('edit', function($scope, $http, $location, editApi) {
    $scope.init = function() {
        $scope.hostName = editApi.getHostName();
        $scope.requests = editApi.getRequests();

        if ($scope.hostName == null || $scope.requests == null) {
            $location.path("/list");
        }
    };

    $scope.editRequest = function() {
        $location.path("/edit/" + $scope.hostName);
        editApi.setEditRequest($scope.requests[this.$index]);
    };

    $scope.init();
}).controller('editRequest', function($scope, $http, $location, editApi) {
    $scope.requests = {};

    $scope.init = function() {
        var request = editApi.getEditRequest();

        if (request == null) {
            $location.path("/list");
            return;
        }

        $scope.hostName = editApi.getHostName();
        $scope.host = request.Host;
        $scope.desc = request.Desc;
        $scope.message = request.Message;
        $scope.method = request.Method;

        //map to array
        var array = [],
            item;
        for (var type in request.Param) {
            if (request.Param.hasOwnProperty(type)) {
                item = {};
                item.key = type;
                item.value = request.Param[type];
                array.push(item);
            }
        }
        $scope.params = array;
        $scope.requests = editApi.getRequests();
    };

    $scope.keyvalueClick = function() {
        var last = $scope.params[$scope.params.length - 1];
        if (last.key != null || last.value != null) {
            $scope.params.push({})
        }
    };

    $scope.delParam = function() {
        var index = this.$index;
        if (~index) $scope.params.splice(index, 1);
    };

    $scope.save = function() {
        $http.post('/api/save', {
            host: $scope.host,
            desc: $scope.desc,
            method: parseInt($scope.method),
            param: $scope.params,
        }).success(function(data) {
            $scope.message = data.message;
            if (data.Code == 200) {

            }
        }).error(function() {
            $scope.message = 'server error : (';
        });
       
    };

    $scope.init();
}).filter("httpMethod", function() {
    return function(type) {
        switch (type) {
            case 0:
                return "GET";
            case 1:
                return "POST";
            default:
                "UNKONW";
        }
    };
})
