<template>
    <v-dialog v-model="allTagsDialog" width = "500" persistent>
        <v-layout justify-center>
             <v-flex>
            <v-card class="pa-3">
                <v-card-title>
                        <h4>All tags:</h4>
                </v-card-title>
                    <v-col v-for= "tag in allTags" :key = "tag.text">
                        <span>@{{tag.text}}</span>
                        <!-- <v-textarea outlined v-model="comment.text" prepend-icon="mdi-comment" rows="2" no-resize readonly></v-textarea> -->
                    </v-col>
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
    name: 'AllTags',
    components: {

    },
    props: {
        allTagsDialog: {
        default: false,
      },
    },
    data() {
        return {
            allTags: [
                {text: "idegas"},
                {text: "idegas"},
                {text: "idegas"},
                {text: "idegas"},
            ]
        }
    },
    mounted() {

    },
    methods: {
        close() {
            this.$emit('update:allTagsDialog', false)
        },
    }
}
</script>