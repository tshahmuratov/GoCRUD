angular.module('bookApp.services',[]).factory('Library',function($resource){
    return $resource('/api/v1/libraries/:id',{id:'@Id'},{
        update: {
            method: 'PUT'
        },
        'get':    {
            url: '/api/v1/libraries/:id',
            method:'GET',
            isArray: true
         },
        addBookToLibrary: {
            url: '/api/v1/libraries/:id',
            method:'POST',
            isArray: false
        }
    });
}).factory('Book',function($resource){
    return $resource('/api/v1/books/:id',{id:'@Id'},{
        update: {
            method: 'PUT'
        },
    });
}).service('popupService',function($window){
    this.showPopup=function(message){
        return $window.confirm(message);
    }
});