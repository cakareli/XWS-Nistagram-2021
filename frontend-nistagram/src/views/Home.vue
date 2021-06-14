<template>
  <v-app class="grey lighten-2">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45"  width="800px" >
            <v-app-bar app height="45" color="grey">
              <v-row>
                <v-col>
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
                <v-spacer></v-spacer>
                <v-col>
                  <v-btn outlined class="mx-5 white">
                    <v-icon>mdi-send</v-icon>
                  </v-btn>
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-toolbar id="guestToolbar" ref="guestToolbar" height="35" class="grey lighten-4" width="800px"  v-show="!loggedUser">
            <v-app-bar app>
              <v-row align="center" justify="space-around">
                <v-col>
                  <v-btn 
                    width="250px"
                    height="35px"
                    @click="$router.push('/registration')"
                    class="grey lighten-2"
                    >Register</v-btn
                  >
                </v-col>
                <v-spacer></v-spacer>
                <v-col>
                  <v-btn
                    width="250px"
                    height="35px"
                    @click="$router.push('/login')"
                    class="grey lighten-2"
                    >Login</v-btn
                  >
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-container>
            <v-row>
              <v-col> </v-col>
            </v-row>
          </v-container>
          <v-bottom-navigation height="35" width="800px" background-color="grey">
            <v-btn value="home">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn value="search">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn value="add" v-show="loggedUser">
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn value="notification" v-show="loggedUser">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            
            <v-btn
              v-show="loggedUser"
              value="profile"
              @click="$router.push('/account')"
            >
              <v-icon>mdi-account</v-icon>
            </v-btn>
              
          </v-bottom-navigation>
        </v-col>
        <v-col width="300px"></v-col>
      </v-row>
    </v-container>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>

import { getId } from '../security/token.js'

export default {
  name: "Home",
  

  data() {
    return {
      loggedUser: false,
    }
  },

  methods: {
    checkLoggedUser(){
      if(getId().length != 0){
        this.loggedUser = true
      }      
    }
  },
  mounted() {
    this.checkLoggedUser();
  },
};

</script>
