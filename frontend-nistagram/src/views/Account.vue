<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-card>
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
              <v-navigation-drawer  v-model="drawer" absolute temporary height="570" class="grey lighten-4">
                <v-list>
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
          </v-card>
            <v-container height="530" class="grey lighten-2">
              <v-row>
                <v-col width="400px"></v-col>
                <v-col>
                  <h3>{{this.username}}</h3>
                </v-col>
              </v-row>
              <v-row>
                <v-col width="400px"></v-col>
                <v-col>
                  <h3>{{this.name}}{{this.surname}}</h3>
                </v-col>
              </v-row>
              <v-row >
                <v-col width="400px"></v-col>
                <v-col>
                  <v-avatar size="80" color="red">
                    <span>{{this.avatar_name}}</span>
                  </v-avatar> 
                  <br>
                  <h3 class="mx-3">Stories</h3> 
                </v-col>
              </v-row> 
              <br><br><br><br><br><br><br><br><br>
              <br><br><br>         
            </v-container>
          <v-bottom-navigation height="35" width="800px" background-color="grey">
            <v-btn value="home" @click="$router.push('/')">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn value="search" @click="$router.push('/search')">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn value="add"  @click="$router.push('/new-post')">
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn value="notification">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            
            <v-btn value="profile">
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
import { getId, getToken, getUsername, removeToken } from '../security/token.js'

export default {
  name: "Account",
  data() {
    return{
      drawer : false,
      name: "",
      surname: "",
      username: "",
      biography: "",
      website: "",
      avatar_name: "",
      imagePaths: [],
      imageIds: [],
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
    },
    findAvatar(){
      this.avatar_name = this.name.substring(0,1) + this.surname.substring(0,1)
    },
    loadUsersPosts(){
      axios.get('http://localhost:8081/api/media-content/regular-user-posts/'+getUsername(),
      {
            headers : {
                          Authorization: 'Bearer ' + getToken()
                      }
      })
      .then(response => {
        this.loadedImages = response.data
        this.imagePaths = this.loadedImages.imagepath
        alert("Image: "+this.imagePaths[0])

      }).catch(error => {
        if(error.response.status === 500){
          console.lof("Internal server error")
      }
    })
    }
  },
  mounted(){
    this.loadRegisteredUser();
    this.findAvatar();
    this.loadUsersPosts();
  }
};
</script>

<style>
</style>