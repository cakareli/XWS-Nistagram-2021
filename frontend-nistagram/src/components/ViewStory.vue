<template>
    <v-dialog v-model="viewStoryDialog" max-width = "550px" max-height="700px" persistent>
        <v-layout justify-center>
             <v-flex>
            <v-card class="pa-3">
                <v-card-title class="grey lighten-3" height="10">
                    <h4>@{{storyView.RegularUser.Username}}</h4>
                    <v-spacer/>
                  </v-card-title>
                  <v-icon class="ml-11">mdi-map-marker</v-icon> {{storyView.Location}}
                  <v-row class="justify-center my-1">
                    <v-img
                      v-bind:src="storyView.MediaPaths[0]"
                      max-width="300"
                      max-height="300"
                      width="400"
                      height="400"
                      class="ma-2"
                    ></v-img>
                  </v-row>
                  <v-row class="justify-center mx-2">
                    <v-textarea
                      label="Description"
                      v-model="storyView.Description"
                      auto-grow
                      outlined
                      rows="1"
                      width="170"
                      readonly
                    ></v-textarea>
                  </v-row>
                  
                  <v-row class="ma-2" justify="center">                   
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllHashtags(storyView.Hashtags)"
                      >Hashtags</v-btn>
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllTags(storyView.Tags)"
                      >Tags</v-btn>
                    <v-btn
                      x-small
                      class="error"
                      @click="reportPost(storyView)"
                    ><v-icon x-small left color="white">mdi-alert-octagon</v-icon>
                      <span>Report</span>
                    </v-btn>
                      <v-spacer/>
                      <v-btn small color="grey lighten-3" @click.native="close">Close</v-btn>
                  </v-row>
            </v-card>
             </v-flex>
        </v-layout>
        <AllTags :allTagsDialog.sync="allTagsDialog" :allPostTags="allPostTags"/>
        <AllHashtags :allHashtagsDialog.sync="allHashtagsDialog" :allPostHashtags="allPostHashtags"/>
        <ReportPost :reportPostDialog.sync="reportPostDialog" :reportedPost="reportedPost"/>
    </v-dialog> 
</template>

<script>

import AllTags from "../components/AllTags.vue";
import AllHashtags from "../components/AllHashtags.vue";
import ReportPost from "../components/ReportPost.vue"

export default {
    name: 'ViewStory',
    components: {
        AllTags,
        AllHashtags,
        ReportPost
    },
    props: {
        viewStoryDialog: {
        default: false,
      },
      storyView: {}
    },
    data() {
        return {
            allPostTags: [],
            allPostHashtags: [], 
            allTagsDialog: false,
            allHashtagsDialog: false,
            reportPostDialog: false,
            reportedPost: {}
        }
    },
    
    methods: {
        close() {
            this.$emit('update:viewStoryDialog', false)
        },
        viewAllTags(allPostTags) {
            this.allPostTags = allPostTags;
            this.allTagsDialog = true;
        },
        viewAllHashtags(allPostHashtags) {
            this.allPostHashtags = allPostHashtags;
            this.allHashtagsDialog = true;
        },
        reportPost(story){
          this.reportPostDialog = true;
          this.reportedPost = story;
        }       
    }
}
</script>