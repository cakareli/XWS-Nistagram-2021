<template>
    <v-app class="grey lighten-3" width="1000px">
        <v-snackbar v-model="snackbar" centered timeout="3500">
            <span>{{snackbarText}}</span>
        </v-snackbar>  
        <v-app-bar app height="45" class="grey lighten-3">
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
        <v-container>
            <v-row justify="center">
                <v-card width="800px" class="pa-12">
                <v-bottom-navigation background-color="grey lighten-3" height="45px">
                    <v-btn @click="postFormButton">
                    <v-icon>mdi-image</v-icon>
                    <span>Post</span>
                    </v-btn>

                    <v-btn @click="albumFormButton">
                    <v-icon>mdi-image-multiple</v-icon>
                    <span>Album</span>
                    </v-btn>

                    <v-btn @click="storyFormButton">
                    <v-icon>mdi-account-supervisor-circle</v-icon>
                    <span>Story</span>
                    </v-btn>
                </v-bottom-navigation>
            </v-card>
            </v-row>
        </v-container>
        <v-container class="justify-center">
            <v-snackbar v-model="snackbar" top timeout="3500">
                <span>{{snackbarText}}</span>
            </v-snackbar>
            <v-container class="pa-12 justify-center" width="1000px" height="1000">
            <v-row>
            <v-col>
                <v-form v-model="form.isFormValid">
                    <v-btn raised class="grey lighten-3" @click="uploadFile">Upload Image</v-btn>
                    <input type="file" style="display: none" ref="fileInput" accept="image/*" @change="onFilePicked"/>
                    <v-text-field label ="Description" v-model="form.description" :rules="rules.description"></v-text-field>
                    <v-text-field label ="Location" v-model="form.location" :rules="rules.location"></v-text-field>
                    <v-row>
                        <v-text-field label ="Hashtags" v-model="hashtag"></v-text-field>
                        <v-btn class="grey lighten-3 ma-3" @click="addHashtag">Add Hashtag</v-btn>
                    </v-row>
                    <br>
                    <v-textarea outlined label ="All hashtags" v-model="form.hashtags" rows="2" no-resize readonly></v-textarea>
                    <v-row>
                        <v-text-field label ="Tags" v-model="tag"></v-text-field>
                        <v-btn class="grey lighten-3 ma-3" @click="addTag">Tag user</v-btn>
                    </v-row>
                    <br>
                    <v-textarea outlined label ="All tags" v-model="form.tags" rows="2" no-resize readonly></v-textarea>
                    <v-btn color="grey lighten-3" allign-right :disabled="!form.isFormValid" @click="submit">{{this.activeForm}}</v-btn>
                </v-form>
            </v-col>
            <v-col v-show="!albumForm">
                <img :src="imageUrl" max-height="400" max-width="400" height="400" width="400"/>
                <v-btn raised class="grey lighten-3" @click="deleteImage">Delete Image</v-btn>
                <v-switch
                    v-show="storyForm"
                    class="mx-15"
                    @click="closeFriendsSwitch"
                    v-model="closeFriends"
                    inset
                    :label="`For close friends: ${closeFriendsString.toString()}`"
                ></v-switch>
            </v-col>
            <v-col v-show="albumForm">
                <v-carousel v-model="model">
                    <v-carousel-item
                    v-for="image in firebaseURL"
                    :key="image"
                    >
                    <v-sheet
                        width="400px"
                        height="400px"
                        tile
                    >
                        <v-row
                        class="fill-height"
                        align="center"
                        justify="center"
                        >
                        <div>
                            <img :src="image" max-height="400" max-width="400" height="400" width="400"/>
                        </div>
                        </v-row>
                    </v-sheet>
                    </v-carousel-item>
                </v-carousel>
            </v-col>
            </v-row>
            </v-container>
        </v-container>

    <v-footer app height="45px" class="grey lighten-3 justify-center">
        <v-container>
          <v-row justify="center">
            <v-btn class= "mx-2" @click="$router.push('/').catch(()=>{})">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn class= "mx-2" @click="$router.push('/search').catch(()=>{})">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn
              class= "mx-2"
              @click="$router.go()"
            >
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn class= "mx-2">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            <v-btn
              class= "mx-2"
              @click="$router.push('/account').catch(()=>{})"
            >
              <v-icon>mdi-account</v-icon>
            </v-btn>
          </v-row>
        </v-container>
    </v-footer>
    </v-app>
</template>

<script>

import firebase from 'firebase/app'
import 'firebase/firebase-storage'
import { getUsername } from '../security/token.js'
import axios from 'axios'

export default {
    name: 'AddPostComment',
    components: {

    },
    data(){
        return {
            snackbar: false,
            snackbarText : "",
            model: 0,
            closeFriends: false,
            closeFriendsString: "No",
            activeForm: "PUBLISH POST",
            postForm: true,
            storyForm: false,
            albumForm: false,
            imageUrl: "",
            image: null,
            uploadValue: 0,
            firebaseURL: [],
            tag : "",
            hashtag: "",
            form:{
                name: "",
                tags: [],
                hashtags: [],
                location: "",
                isFormValid: false,
            },
            rules:{
                descriptionRules: [
                    description => !!description || 'Description is required!',
                    description => (description && description.length <= 100) || 'Description must contain less than 100 characters!',
                ],
                locationRules: [
                    description => !!description || 'Description is required!',
                    description => (description && description.length <= 30) || 'Description must contain less than 100 characters!',
                ],
            },
        }
    },
    mounted() {

    },
    methods: {
        closeFriendsSwitch(){
            if(this.closeFriends){
                this.closeFriendsString = "Yes"
            }
            else{
                this.closeFriendsString = "No"
            }
        },
        postFormButton(){
            this.postForm = true;
            this.storyForm = false;
            this.albumForm = false
            this.activeForm = "PUBLISH POST";
        },
        storyFormButton(){
            this.postForm = false;
            this.storyForm = true;
            this.albumForm = false
            this.activeForm = "PUBLISH STORY";
        },
        albumFormButton(){
            this.postForm = false;
            this.storyForm = false;
            this.albumForm = true
            this.activeForm = "PUBLISH ALBUM";
        },
        deleteImage(){
            if(this.firebaseURL.length >= 0 && (this.postForm==true || this.storyForm == true)){
                this.firebaseURL[0] = ""
                this.imageUrl = ""
            }else{
                this.snackbar = true
                this.snackbarText = "There is no uploaded image"
            }           
        },
        uploadFile(){
            if(this.firebaseURL.length>=1 && (this.postForm==true || this.storyForm == true)){
                this.snackbar = true
                this.snackbarText = "Delete image first"
            }else{
                this.$refs.fileInput.click();
            }
        },
        onFilePicked(event){    
            const files = event.target.files;
            let filename = files[0].name;
            console.log(filename)
            if(filename.lastIndexOf('.') <= 0){
               return alert("Please upload valid image file!")
            }

            const fileReader = new FileReader();
            fileReader.addEventListener('load', () =>{
                this.imageUrl = fileReader.result;
            })
            fileReader.readAsDataURL(files[0])
            this.image = files[0]

            const storageRef = firebase.storage().ref(`${filename}`).put(this.image)
            storageRef.on(`state_changed`,snapshot=>{
                this.uploadValue = (snapshot.bytesTransferred/snapshot.totalBytes)*100;
                }, error=>{console.log(error.message)},
                ()=>{this.uploadValue=100;
                storageRef.snapshot.ref.getDownloadURL().then((url)=>{
                this.firebaseURL.push(url)
                console.log(this.firebaseURL)
                });
            });
        },
        addTag(){
            if(this.tag != "" && !this.form.tags.includes(this.tag))
                this.form.tags.push(this.tag);
            this.tag = "";
        },
        addHashtag(){
            if(this.hashtag != "" && !this.form.hashtags.includes(this.hashtag))
                this.form.hashtags.push(this.hashtag);
            this.hashtag = "";
        },
        submit() {
            if(this.firebaseURL == ""){
                alert("Image is still not uploaded!")
            }else {

                if(!this.storyForm){
                   
                    let postUploadDTO = {
                hashtags : this.form.hashtags,
                tags : this.form.tags,
                description : this.form.description,
                mediaPaths : this.firebaseURL,
                uploadDate: new Date(),
                location: this.form.location,
                username: getUsername(),
                }
                axios.post('http://localhost:8081/api/media-content/new-post',
                    postUploadDTO
                ).then(response => {
                    console.log(response)
                    this.$router.push('/').catch(()=>{})
                }).catch(error => {
                    if(error.response.status === 500){
                        this.snackbarText = "Internal server error occurred!";
                        this.snackbar = true;
                    }else if(error.response.status === 400){
                        this.snackbarText = "Bad request, try again!";
                        this.snackbar = true;
                    }
                })
                }
                else if(this.storyForm){
                        let storyUploadDTO = {
                    hashtags : this.form.hashtags,
                    tags : this.form.tags,
                    description : this.form.description,
                    mediaPaths : this.firebaseURL,
                    uploadDate: new Date(),
                    location: this.form.location,
                    username: getUsername(),
                    forCloseFriends: this.closeFriends
                    }
                    axios.post('http://localhost:8081/api/media-content/new-story',
                        storyUploadDTO
                    ).then(response => {
                        console.log(response)
                        this.$router.push('/').catch(()=>{})
                    }).catch(error => {
                        if(error.response.status === 500){
                            this.snackbarText = "Internal server error occurred!";
                            this.snackbar = true;
                        }else if(error.response.status === 400){
                            this.snackbarText = "Bad request, try again!";
                            this.snackbar = true;
                        }
                    })
                }                
            }
        }
    }
}
</script>
