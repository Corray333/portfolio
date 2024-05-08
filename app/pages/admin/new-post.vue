<script setup>
import { ref } from 'vue'
import axios from 'axios'
import markdownit from 'markdown-it'

const md = markdownit({
    html: true,
    linkify: true,
    typographer: true,
    xhtmlOut: true,
})

const post = ref({
    title: "",
    description: "",
    content: "",
    tags: [],
    cover:""
})

const tag = ref("")

const newTag = () => {
    post.value.tags.push(tag.value)
    tag.value = ""
}

const file = ref(null)

const handleCoverUpload =  async (event) => {
    if (event.target.files[0].size > 5 * 1024*1024) {
        return
    }
    file.value = event.target.files[0]
    const reader = new FileReader()

    reader.onload = (e) => {
        post.value.cover = e.target.result
    }
    reader.readAsDataURL(event.target.files[0])
    post.value.cover = await loadFile()
}

const loadFile = async ()=>{
    try{
        let formData = new FormData()
        formData.append('file', file.value)

        let {data} = await axios.post(`/api/upload/image`, formData, {
            headers:{
                "Authorization":localStorage.getItem('Authorization')
            }
        })
        return data.url
    } catch(err){
        console.log(err)
    }
}

const handleFileUpload =  async (event) => {
    if (event.target.files[0].size > 5 * 1024*1024) {
        return
    }
    file.value = event.target.files[0]
    const reader = new FileReader()

    reader.readAsDataURL(event.target.files[0])
    let file_name = await loadFile()
    navigator.clipboard.writeText(`![image](${file_name})`)
}   


</script>

<template>
    <section class="editor w-full min-h-screen z-50 relative pointer p-20 flex flex-col items-center">
        <h1 class="text-center text-[4rem]">New post</h1>
        <section class="post-container max-w-[80rem] w-full flex flex-col gap-2">
            <div class="profile_photo relative">
                <input @input="changed = true" type="file" id="fileInput" class="hidden" @change="handleCoverUpload" />
                <label for="fileInput"
                    class="text-center absolute mx-auto bg-gray-900 bg-opacity-80 h-full w-full flex items-center justify-center opacity-0 duration-300 cursor-pointer  hover:opacity-100">
                    <Icon icon="mdi:user" />
                </label>
                <img :src="file ? post.cover : '/api/images/noimage.webp'" alt=""
                    class="w-full h-48 object-cover border-white">
            </div>
            <section class="header flex gap-10 py-5">
                <div class="flex flex-col gap-5 w-full">
                    <div class="flex gap-2">
                        <label for="">Title:</label>
                        <input type="text" placeholder="New post" class=" border-b-2 border-dark">
                    </div>
                    <div>
                        <div class="flex gap-2">
                            <label for="">Tags:</label>
                            <div>
                                <input @keyup.enter="newTag" v-model="tag" type="text"
                                placeholder="golang, lifestyle, etc. " class=" border-b-2 border-dark">
                                <div class="flex gap-2">
                                    <p><br></p>
                                    <p v-for="(tag, i) of post.tags" :key="i" @click="post.tags.splice(i, 1)"
                                        class=" underline cursor-pointer">{{ tag }}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="">Description:</label>
                    <textarea type="text" placeholder="Post description"
                        class="w-full border-l-2 border-dark pl-2 h-full"></textarea>
                </div>
            </section>
            <div class="flex">
                <button class=" duration-300 hover:bg-dark hover:text-white p-2 rounded-full"><Icon name="mdi:camera" class="text-2xl" /></button>
            </div>
            <section class="content grid grid-cols-2 gap-10">
                <div class="editor flex flex-col gap-2">
                    <textarea v-model="post.content" name="" id="" cols="30" rows="10"
                        class="content-input p-5 w-full border-2 border-dashed border-dark"></textarea>
                </div>
                <article class="content p-5 viewport w-full border-2 border-dashed border-dark"
                    v-html="md.render(post.content)"></article>
            </section>
        </section>
    </section>
</template>


<style scoped>
.editor {
    cursor: default !important;
}
</style>