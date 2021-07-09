<template>
  <v-app class="grey lighten-4" width="800px">
    <v-app-bar app height="45" color="grey lighten-3">
      <v-app-bar-nav-icon @click="$router.push('/administratorHome')">
                <v-icon>mdi-arrow-left</v-icon>
        </v-app-bar-nav-icon>
        <v-row>
          <v-col>
            <v-toolbar-title>
              <span outlined class="font-weight-light">NISTA</span>
              <span>GRAM</span>
            </v-toolbar-title>
          </v-col>
          <v-spacer></v-spacer>
        </v-row>
    </v-app-bar>

    <v-row justify="center" class="pa-16">
            <v-list>
              <v-list-item v-for="user in allUsers" :key="user._id">
                <v-card height="60" width="600" class="ma-2 grey lighten-5">
                  <v-card-title height="60" class="grey lighten-3">
                    <v-row>
                      <v-col>
                        <v-row class="pa-3">
                          <h4>@{{ user.username }}</h4>
                        </v-row>
                        
                      </v-col>
                      <v-spacer/>
                      <v-col>
                        <v-btn small @click="loadUserProfile(user.username)">
                          <span>View profile</span>
                        </v-btn>
                      </v-col>
                      <v-col>
                        <v-btn small class="error" @click="removeUser(user._id)">
                          <v-icon left>mdi-trash-can-outline</v-icon>
                          <span>Remove</span>
                        </v-btn>
                      </v-col>
                    </v-row>
                  </v-card-title>
                  <v-divider></v-divider>
                </v-card>
              </v-list-item>
            </v-list>
          </v-row>

    <v-footer app height="45px" class="grey lighten-3 justify-center">
        <v-container>
          <v-row justify="center">      
            <v-btn class= "mx-2" @click="$router.push('/administratorHome').catch(()=>{})">
              <v-icon>mdi-home</v-icon>
            </v-btn>
          </v-row>
        </v-container>
    </v-footer>

  </v-app>
</template>

<script>

import axios from "axios";
import { getToken} from "../security/token.js";

export default {
    name: 'RemoveProfiles',
    data() {
        return{
            allUsers: []
        }
    },
    methods: {
        loadAllRegularUser(){
            axios.get("http://localhost:8081/api/user/get-all-regular-users",{
                headers: {
                    Authorization: "Bearer " + getToken(),
                },
            })
            .then((response)=>{
                this.allUsers = response.data
            })
            .catch((error)=>{
                if(error.response===500){
                    console.log("Internal server error")
                }
            })
        },
        loadUserProfile(username){
            this.$router.push('/user-profile/' + username).catch(()=>{})
        },
        removeUser(id){
          axios.delete('http://localhost:8081/api/user/delete-regular-user/'+id,{
            headers: {
                    Authorization: "Bearer " + getToken(),
                },
          })
          .then((response)=>{
            console.log("User removed")
            console.log(response.data.status)
            this.$router.go()
          })
          .catch((error)=>{
            if(error.response.data.status === 404){
              console.log("Status not found")
            }
          })
        }
    },
    mounted() {
        this.loadAllRegularUser();
    }
}
</script>

<style>

</style>