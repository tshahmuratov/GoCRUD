angular.module('bookApp.services',[]).factory('Book',function($resource){
    return $resource('/api/v1/books/:id',{id:'@_id'},{
        update: {
            method: 'PUT'
        }
    });
}).service('popupService',function($window){
    this.showPopup=function(message){
        return $window.confirm(message);
    }
});