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
                        <v-btn small class="success" @click="verifyUser(user._id)">
                            <v-icon left>mdi-check-circle-outline</v-icon>
                            <span>Verify</span>
                      </v-btn>
                      </v-col>
                      <v-col>
                        <v-btn small class="error" @click="removeUser(user._id)">
                          <v-icon left>mdi-trash-can-outline</v-icon>
                          <span>Reject</span>
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
    name: 'AgentVerification',
    data() {
        return {

        }
    },
    methods: {
        loadAllAgentRequest(){
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
    },
    mounted() {
        this.loadAllAgentRequest()
    }
}
</script>

<style>

</style>