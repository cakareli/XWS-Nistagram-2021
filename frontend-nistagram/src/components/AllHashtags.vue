<template>
    <v-dialog v-model="allHashtagsDialog" width = "500" persistent>
        <v-layout justify-center>
             <v-flex>
            <v-card class="pa-3">
                <v-card-title>
                        <h4>All hashtags:</h4>
                </v-card-title>
                    <span v-for="(hashtag, index) in allPostHashtags" :key ="index" class="ma-2">#{{hashtag}} </span>
                <v-card-actions>
                    <v-btn color="grey lighten-3" @click.native="close">Close</v-btn>
                </v-card-actions>
            </v-card>
             </v-flex>
        </v-layout>
         
    </v-dialog>
</template>
    
    
<script>

export default {
    name: 'AllHashtags',
    components: {

    },
    props: {
        allHashtagsDialog: {
        default: false,
      },
      allPostHashtags: {}
    },
    data() {
        return {

        }
    },
    mounted() {

    },
    methods: {
        close() {
            this.$emit('update:allHashtagsDialog', false)
        },
    }
}
</script>