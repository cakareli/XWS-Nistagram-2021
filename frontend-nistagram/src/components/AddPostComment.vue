<template>
    <v-dialog v-model="addPostCommentDialog" width = "500" persistent>
        <v-layout justify-center>
             <v-flex>
            <v-card class="pa-3">
                <v-card-title>
                        <h4>Add new comment:</h4>
                </v-card-title>
                    <v-form v-model="form.isFormValid">
                        <v-textarea outlined label ="Comment text" v-model="form.text" :rules="rules.textRules" prepend-icon="mdi-comment" rows="3" no-resize></v-textarea>
                    </v-form>
                <v-card-actions>
                    <v-btn color="grey lighten-3" allign-right :disabled="!form.isFormValid" @click="submit">Add</v-btn>
                     <v-spacer/>
                    <v-btn color="grey lighten-3" @click.native="close">Close</v-btn>
                </v-card-actions>
            </v-card>
             </v-flex>
        </v-layout>
    </v-dialog>
</template>
    
    
<script>

import axios from 'axios'
import { getToken, getUsername } from '../security/token.js'

export default {
    name: 'AddPostComment',
    components: {

    },
    props: {
        addPostCommentDialog: {
        default: false,
        },
        postId: {}
    },
    data() {
        return {
            form: {
                text: "",
            },
            rules:{ 
                textRules: [
                    text => !!text || 'Comment text is required!',
                ],
            },
        }
    },
    mounted() {

    },
    methods: {
        close() {
            this.$emit('update:addPostCommentDialog', false)
        },
        submit() {
            let commentPostDTO = {
                username: getUsername(),
                postId: this.postId,
                text: this.form.text
            }
        axios.put("http://localhost:8081/api/media-content/comment-post",
            commentPostDTO,
        {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
            console.log(response)
            this.$router.go()
        });
        }
    }
}
</script>