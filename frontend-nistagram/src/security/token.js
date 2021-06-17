import VueJwtDecode from 'vue-jwt-decode'

function setToken(token){
    localStorage.setItem('jwt', token)
}

function getToken(){
   return localStorage.getItem('jwt')
}

function getRole(){
    var decoded = VueJwtDecode.decode(getToken())
    return decoded.role
}

function getUsername(){
    var decoded = VueJwtDecode.decode(getToken())
    return decoded.username
}

function getId() {
    var decoded = VueJwtDecode.decode(getToken())
    return decoded.userId
}

function removeToken(){
    localStorage.removeItem('jwt')
}

export{
    setToken,
    getToken,
    getRole,
    getId,
    getUsername,
    removeToken,
}