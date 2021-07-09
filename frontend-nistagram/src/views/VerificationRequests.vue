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

    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">         
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="verificationRequest in allVerficationRequests" :key="verificationRequest._id">
                <v-card height="600" width="550" class="ma-3 grey lighten-5">
                  <v-card-title class="grey lighten-3" height="10">
                      <h4>{{verificationRequest.verificationName}} {{verificationRequest.verificationSurname}}</h4>
                      <v-spacer></v-spacer>
                      <h4>{{verificationRequest.verificationType}}</h4>
                  </v-card-title>
                  <v-row class="justify-center my-1">
                    <v-img
                      v-bind:src="verificationRequest.verificationImageUrl"
                      max-width="400"
                      max-height="400"
                      width="400"
                      height="400"
                      class="ma-2"
                    ></v-img>
                  </v-row>
                  <v-row justify="center">
                      <v-btn width="200" class="mx-5 success" @click="verifyUser(verificationRequest.userId, verificationRequest.verificationType,verificationRequest._id)">
                        <v-icon left>mdi-check-circle-outline</v-icon>
                        <span>Accept</span>
                      </v-btn>                     
                      <v-btn width="200" class="mx-5 error" @click="rejectVerification(verificationRequest._id)">
                        <v-icon left>mdi-trash-can-outline</v-icon>
                        <span>Reject</span>
                      </v-btn>
                  </v-row>
                </v-card>
                <v-divider></v-divider>
              </v-list-item>
            </v-list>
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

import axios from "axios";
import { getToken} from "../security/token.js";

export default {
    name: 'VerificationRequests',
    data() {
        return {
            allVerficationRequests: [],
            verificationType: ""
        }
    },
    methods: {
        loadAllVerificationRequests(){
            axios.get("http://localhost:8081/api/user/get-all-verification-requests",{
                headers: {
                    Authorization: "Bearer " + getToken(),
                },
            }).then((response)=>{
                this.allVerficationRequests = response.data
                console.log(response.status)
                for( let i = 0; i < this.allVerficationRequests.length; i++){
                    if(this.allVerficationRequests[i].verificationType === 1){
                        this.allVerficationRequests[i].verificationType = "Influencer"
                    }
                    else if(this.allVerficationRequests[i].verificationType === 2){
                        this.allVerficationRequests[i].verificationType = "Sports"
                    }
                    else if(this.allVerficationRequests[i].verificationType === 3){
                        this.allVerficationRequests[i].verificationType = "Business"
                    }
                    else if(this.allVerficationRequests[i].verificationType === 4){
                        this.allVerficationRequests[i].verificationType = "Brand"
                    }
                    else if(this.allVerficationRequests[i].verificationType === 5){
                        this.allVerficationRequests[i].verificationType = "Organization"
                    }
                    else if(this.allVerficationRequests[i].verificationType === 6){
                        this.allVerficationRequests[i].verificationType = "NewMedia"
                    }
                }
            })
            .catch((error)=>{
                if(error.status === 500){
                    console.log("Internal server error")
                }
            })
        },
        verifyUser(id, type, verificationId){
          if(type === "Influencer"){
               type = 1
          }
          else if(type === "Sports"){
               type = 2
          }
          else if(type === "Business"){
               type = 3
          }
          else if(type === "Brand"){
              type = 4
          }
          else if(type === "Organization"){
              type = 5
          }
          else if(type === "NewMedia"){
              type = 6
          }
          let userVerification = {
            _id : id,
            verificationType : type
          }
          alert("Type: "+type)
          axios.post('http://localhost:8081/api/user/verify-user',userVerification ,{
            headers: {
                    Authorization: "Bearer " + getToken(),
                },
          })
          .then((response)=>{
            console.log(response.status)
            this.rejectVerification(verificationId)
            this.$router.go()
          })
          .catch((error)=>{
            if(error.status === 404){
              console.log("Bad request")
            }
          })
        },
        rejectVerification(id){
          axios.delete('http://localhost:8081/api/user/delete-verification-request/'+id, {
            headers: {
                    Authorization: "Bearer " + getToken(),
                },
          })
          .then((response)=>{
            console.log(response.status)
            this.$router.go()
          })
          .catch((error)=>{
            if(error.response.status === 404){
              console.log("Status bad request")
            }
          })
        }
    },
    mounted() {
        this.loadAllVerificationRequests();
    }

}
</script>

<style>

</style>