<template>
  <v-app class="grey lighten-4" width="800px">

    <v-app-bar app height="45" color="grey lighten-3">
      <v-app-bar-nav-icon @click.stop = "drawer=!drawer" class="ml-5"/>
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

    <v-navigation-drawer
          v-model="drawer"
          class="grey lighten-4 pa-8"
          absolute
        >
          <v-list>
            <v-list-item-group active-class="deep-purple--text text--accent-4">

              <v-list-item @click="logout">
                <v-list-item-title>
                  <v-icon small>mdi-logout</v-icon>
                  Log out</v-list-item-title
                >
              </v-list-item>
              
            </v-list-item-group>
          </v-list>
    </v-navigation-drawer>

    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12 grey lighten-4">
            <v-row justify="center">
                
                <v-col>
                    <br>
                    <h3>{{this.name}} {{this.space}} {{this.surname}}</h3>
                    <br>
                    
                </v-col>
            </v-row>
        </v-card>
      </v-row> 
    </v-container>
    <br><br><br>
    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12 grey lighten-4">
            <v-row justify="center" class="ma-5">
                <v-btn width="250" height="35" @click="$router.push('/verificationRequests')">Verification Requests</v-btn>
            </v-row>
            <v-row justify="center" class="ma-5">
                <v-btn width="250" height="35" @click="$router.push('/inappropriateContent')">Inappropriate Content</v-btn> 
            </v-row>
            <v-row justify="center" class="ma-5">
                <v-btn width="250" height="35" @click="$router.push('/removeProfiles')">Remove Profile</v-btn> 
            </v-row>
        </v-card>
      </v-row> 
    </v-container>    

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

import {removeToken } from "../security/token.js";

export default {
    name: 'AdministratorHome',
    data() {
        return {
            drawer: false
        }
    },
    methods: {
        logout() {
            removeToken();
            this.$router.push("/").catch(() => {});
        },
    }
    
}
</script>

<style>

</style>