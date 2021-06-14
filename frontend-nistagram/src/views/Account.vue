<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45" width="800px" >
            <v-app-bar  app height="45" color="grey">
              <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
              <v-row >
                <v-col>
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
                <v-spacer></v-spacer>
                <v-col>
                  <v-btn outlined class="mx-5 white" small>
                    <v-icon>mdi-send</v-icon>
                  </v-btn>
                </v-col>
              </v-row>
            </v-app-bar>
            <v-navigation-drawer v-model="drawer" absolute temporary height="500" class="grey lighten-4">
              <v-list nav dense>
                <v-list-item-group
                  active-class="deep-purple--text text--accent-4">
                  <v-list-item>
                    <v-list-item-title>
                      <v-icon small>mdi-clock</v-icon>
                      Activity
                    </v-list-item-title>
                  </v-list-item>

                  <v-list-item @click="$router.push('/update')">
                    <v-list-item-title>
                      <v-icon small>mdi-pencil</v-icon>
                      Edit Profile
                      </v-list-item-title>
                  </v-list-item>

                  <v-list-item>
                    <v-list-item-title>
                      <v-icon small>mdi-lock</v-icon>
                      Privacy</v-list-item-title>
                  </v-list-item>

                  <v-list-item>
                    <v-list-item-title>
                      <v-icon small>mdi-bookmark</v-icon>
                      Saved Posts</v-list-item-title>
                  </v-list-item>

                  <v-list-item>
                    <v-list-item-title>
                      <v-icon small>mdi-text-box-multiple</v-icon>
                      Verify</v-list-item-title>
                  </v-list-item>

                  <v-list-item @click="logout">
                    <v-list-item-title>
                      <v-icon small>mdi-logout</v-icon>
                      Log out</v-list-item-title>
                  </v-list-item>

                </v-list-item-group>
              </v-list>
            </v-navigation-drawer>
          </v-toolbar>

          <v-container>
            <v-row height="100px" width="600px">
              <v-col width="50px">
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <label>{{this.name}}</label>
                <label>{{this.surname}}</label>
              </v-col>
              <v-col>
                <v-btn height="400">

                </v-btn>
              </v-col>
            </v-row>
          </v-container>
          <v-bottom-navigation height="35" width="800px" background-color="grey">
            <v-btn value="home" @click="$router.push('/')">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn value="search">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn value="add">
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn value="notification">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            
            <v-btn
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
  </v-app>
</template>

<script>

import axios from "axios";
import { getId, getToken, removeToken } from '../security/token.js'

export default {
  name: "Account",
  data() {
    return{
      drawer : false,
      name: "",
      surname: "",
      username: "",
      biography: "",
      website: ""
    }
  },
  methods: {
    loadRegisteredUser(){
            axios.get('http://localhost:8081/api/user/logged-user/'+getId(),
            {
                    headers : {
                        Authorization: 'Bearer ' + getToken()
                    }
            })
            .then(response => {
                this.regularUser = response.data;
                this.name = this.regularUser.Name + " ";
                this.surname = this.regularUser.Surname;
                this.username = this.regularUser.Username;
                this.biography = this.regularUser.Biography;
                this.website = this.regularUser.WebSite;
            }).catch(error => {
            if(error.response.status === 500){
                this.snackbarText = "Internal server error occurred!";
                this.snackbar = true;
            }
            if(error.response.status === 401){
               this.snackbar = true
               this.snackbarText = "You are unauthorized to get patient informations!";
            }
        })
    },
    logout(){
      removeToken()
        this.$router.push("/").catch(()=>{})
    }
  },
  mounted(){
    this.loadRegisteredUser();
  }
};
</script>

<style>
</style>