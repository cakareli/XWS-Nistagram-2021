<template>
    <v-app>
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
                <v-bottom-navigation background-color="grey lighten-3" height="45px">
                    <v-btn @click="postFormButton">
                    <v-icon>mdi-image</v-icon>
                    <span>Post</span>
                    </v-btn>

                    <v-btn @click="storyFormButton">
                    <v-icon>mdi-account-supervisor-circle</v-icon>
                    <span>Story</span>
                    </v-btn>
                </v-bottom-navigation>
            </v-card>
            </v-row>
        </v-container>

        <v-container v-show="!empty">
      <v-row justify="center">
        <v-card width="800px" class="pa-12">         
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="reported in allReports" :key="reported.id">
                <v-card height="750" width="550" class="ma-3 grey lighten-5">
                  <v-card-title class="grey lighten-3" height="10">
                    <a :href="'/user-profile/' + reported.content.RegularUser.Username" class="black--text" style="text-decoration: none; color: inherit;">@{{ reported.content.RegularUser.Username }}</a>
                    <v-spacer/>
                  </v-card-title>
                  <v-icon class="ml-11">mdi-map-marker</v-icon> {{reported.content.Location}}
                  <v-row class="justify-center my-1" v-show="reported.content.MediaContentType==0 || reported.content.MediaContentType==1">
                    <v-img
                      v-bind:src="reported.content.MediaPaths[0]"
                      max-width="400"
                      max-height="400"
                      width="400"
                      height="400"
                      class="ma-2"
                    ></v-img>
                  </v-row>
                  <v-row class="justify-center my-1" v-show="reported.content.MediaContentType==2">
                    <v-carousel class="mx-16" >
                      <v-carousel-item
                      height="300px"
                      v-for="image in reported.content.MediaPaths"
                      :key="image"
                      :src="image">                     
                      </v-carousel-item>
                    </v-carousel>
                  </v-row>
                  <v-row class="justify-center mx-2">
                    <v-textarea
                      label="Description"
                      v-model="reported.content.Description"
                      auto-grow
                      outlined
                      rows="1"
                      width="170"
                      readonly
                    ></v-textarea>
                  </v-row>
                  <v-row class="ma-2" height="5" v-show="postForm">
                    <span> Likes: {{ reported.content.Likes }}</span>
                    <v-spacer />
                    <span> Dislikes: {{ reported.content.Dislikes }}</span>
                  </v-row>
                  <br>
                  <v-row class="ma-2">
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllHashtags(reported.content.Hashtags)"
                      >Hashtags</v-btn
                    >
                    <v-spacer></v-spacer>
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllTags(reported.content.Tags)"
                      >Tags</v-btn
                    >
                    <v-spacer></v-spacer>
                    
                    <v-btn x-small @click="viewAllPostComments(reported.content.Comment)"
                      >View all comments</v-btn
                    >
                  </v-row>
                  <v-row justify="center">
                      <v-btn width="200" class="mx-5 success" @click="rejectReport(reported._id)">
                        <v-icon left>mdi-check-circle-outline</v-icon>
                        <span>Reject report</span>
                      </v-btn>                     
                      <v-btn width="200" class="mx-5 error" @click="deletePost(reported.content.Id, reported._id)">
                        <v-icon left>mdi-trash-can-outline</v-icon>
                        <span>Delete Content</span>
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
import {getToken} from "../security/token.js";

export default {
    name: 'InappropriateContent',
    data() {
        return{
            postForm: true,
            storyForm: false,
            allReports: [],
            empty: false
        }
    },
    methods: {
        postFormButton(){
            this.postForm = true;
            this.storyForm = false;
            axios.get('http://localhost:8081/api/media-content/get-all-report-post',
                {
                    headers: {
                    Authorization: "Bearer " + getToken(),
                    },
                })
                .then((response)=>{
                    this.empty = false
                    this.allReports = response.data
                    if(this.allReports.length === 0){
                      this.empty = true
                    }
                    console.log(response.status)
                })
                .catch((error)=>{
                    if(error.response.status === 505)
                        console.log("Internal server error")
            })
        },
        storyFormButton(){
            this.postForm = false;
            this.storyForm = true;
            axios.get('http://localhost:8081/api/media-content/get-all-report-story',
                {
                    headers: {
                    Authorization: "Bearer " + getToken(),
                    },
                })
                .then((response)=>{
                    this.empty = false
                    this.allReports = response.data
                    if(response.data === null){
                      this.empty = true
                    }
                    console.log(response.status)
                })
                .catch((error)=>{
                    if(error.response.status === 505)
                        console.log("Internal server error")
            })
        },
        rejectReport(reportId){
          if(this.postForm){
            axios.delete('http://localhost:8081/api/media-content/delete-reported-post/'+reportId,{
              headers: {
                    Authorization: "Bearer " + getToken(),
              },
            })
            .then((response)=>{
              console.log(response.data.status)
              console.log("Post report deleted")
              this.$router.go()
            })
            .catch((error)=>{
              if(error.response.status === 404){
                console.log("Status not found")
              }
            })
          }else{
            axios.delete('http://localhost:8081/api/media-content/delete-reported-story/'+reportId,{
              headers: {
                    Authorization: "Bearer " + getToken(),
              },
            })
            .then((response)=>{
              console.log(response.data.status)
              console.log("Post report deleted")
              this.$router.go()
            })
            .catch((error)=>{
              if(error.response.status === 404){
                console.log("Status not found")
              }
            })
          }
        },
        deletePost(contentId, reportId){
           if(this.postForm){
            axios.delete('http://localhost:8081/api/media-content/delete-post/'+contentId,{
              headers: {
                    Authorization: "Bearer " + getToken(),
              },
            })
            .then((response)=>{
              console.log(response.data.status)
              console.log("Post report deleted")
              this.rejectReport(reportId)
              this.$router.go()
            })
            .catch((error)=>{
              if(error.response.status === 404){
                console.log("Status not found")
              }
            })
          }else{
            axios.delete('http://localhost:8081/api/media-content/delete-story/'+contentId,{
              headers: {
                    Authorization: "Bearer " + getToken(),
              },
            })
            .then((response)=>{
              console.log(response.data.status)
              console.log("Post report deleted")
              this.rejectReport(reportId)
              this.$router.go()
            })
            .catch((error)=>{
              if(error.response.status === 404){
                console.log("Status not found")
              }
            })
          }
        }
    },
    mounted() {
        this.postFormButton();
    }
}
</script>

<style>

</style>