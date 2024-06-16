<script setup>
import { ref } from 'vue'
import axios from 'axios'
import markdownit from 'markdown-it'
import {useRouter} from 'vue-router'

const router = useRouter()

const md = markdownit({
    html: true,
    linkify: true,
    typographer: true,
    xhtmlOut: true,
})

const posts = ref([{
    title: "",
    description: "",
    content: "",
    tags: [],
    cover: "",
    lang: "ru",
}])

const langs = ref(["ru"])
const langPick = ref(0)
const new_lang = ref("")

const newLang = () => {
    posts.value.push({
        title: "",
        description: "",
        content: "",
        tags: [],
        cover: "",
        lang: new_lang.value,
    })
    langs.value.push(new_lang.value)
    new_lang.value = ""
}

const tag = ref("")

const newTag = () => {
    for (let i = 0; i < posts.value.length; i++) {
        posts.value[i].tags.push(tag.value)
    }
    tag.value = ""
}

const removeTag = (i) => {
    for (let i = 0; i < posts.value.length; i++) {
        posts.value[i].tags.splice(i, 1)
    }
}

const file = ref(null)

const handleCoverUpload = async (event) => {
    if (event.target.files[0].size > 5 * 1024 * 1024) {
        return
    }
    file.value = event.target.files[0]
    const reader = new FileReader()

    reader.onload = (e) => {
        posts.value[0].cover = e.target.result
    }
    reader.readAsDataURL(event.target.files[0])
    let url =  await loadFile()
    for (let i = 0; i<posts.value.length; i++){
        posts.value[i].cover = url
    }
}

const loadFile = async () => {
    try {
        let formData = new FormData()
        formData.append('file', file.value)

        let { data } = await axios.post(`/api/upload/image`, formData, {
            headers: {
                "Authorization": localStorage.getItem('Authorization')
            }
        })
        return data.url
    } catch (err) {
        console.log(err)
    }
}

const handleImageUpload = async (event) => {
    if (event.target.files[0].size > 5 * 1024 * 1024) {
        return
    }
    file.value = event.target.files[0]
    const reader = new FileReader()

    reader.readAsDataURL(event.target.files[0])
    let file_name = await loadFile()
    console.log(`![image](${file_name})`)
    navigator.clipboard.writeText(`![image](${file_name})`)
}

const createPost = async ()=>{
    try{
        let {data} =  await axios.post(`/api/posts`, posts.value,{
            headers:{
                Authorization:localStorage.getItem('Authorization')
            }
        })
        router.push(`/blog/${data.id}`)
    } catch(err){
        console.log(err)
    }
}


</script>

<template>
    <section class="editor w-full min-h-screen z-50 relative pointer p-20 gap-5 flex flex-col items-center">
        <h1 class="text-center text-[4rem]">New post</h1>
        <section class="post-container max-w-[80rem] w-full flex flex-col gap-2">
            <div class="profile_photo relative">
                <input @input="changed = true" type="file" id="coverInput" class="hidden" @change="handleCoverUpload" />
                <label for="coverInput"
                    class="text-center absolute mx-auto bg-gray-900 bg-opacity-80 h-full w-full flex items-center justify-center opacity-0 duration-300 cursor-pointer  hover:opacity-100">
                    <Icon icon="mdi:user" />
                </label>
                <img :src="file ? posts[0].cover : '/api/images/noimage.webp'" alt=""
                    class="w-full h-48 object-cover border-white">
            </div>
            <section class="header flex gap-10 py-5">
                <div class="flex flex-col gap-5 w-full">
                    <div class="flex gap-2">
                        <label for="">Languages:</label>
                        <div>
                            <input @keyup.enter="newLang" v-model="new_lang" type="text" placeholder="ru, eng, etc. "
                                class=" border-b-2 border-dark">
                            <div class="flex gap-2">
                                <p><br></p>
                                <p v-for="(lang, i) of langs" :key="i" @click="langPick = i"
                                    class=" underline cursor-pointer" :class="langPick == i ? 'font-bold' : ''">{{ lang
                                    }}
                                </p>
                            </div>
                        </div>
                    </div>
                    <div class="flex gap-2">
                        <label for="">Title:</label>
                        <input v-model="posts[langPick].title" type="text" placeholder="New post" class=" border-b-2 border-dark">
                    </div>
                    <div>
                        <div class="flex gap-2">
                            <label for="">Tags:</label>
                            <div>
                                <input @keyup.enter="newTag" v-model="tag" type="text"
                                    placeholder="golang, lifestyle, etc. " class=" border-b-2 border-dark">
                                <div class="flex gap-2">
                                    <p><br></p>
                                    <p v-for="(tag, i) of posts[0].tags" :key="i" @click="removeTag(i)"
                                        class=" underline cursor-pointer">{{ tag }}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="">Description:</label>
                    <textarea v-model="posts[langPick].description" type="text" placeholder="Post description"
                        class="w-full border-l-2 border-dark pl-2 h-full"></textarea>
                </div>
            </section>
            <div class="flex">
                <span>
                    <input @input="changed = true" type="file" id="imageInput" class="hidden"
                        @change="handleImageUpload" />
                    <label for="imageInput"
                        class=" duration-300 hover:bg-dark hover:text-white p-2 rounded-full w-12 h-12 flex justify-center items-center cursor-pointer">
                        <Icon name="mdi:camera" class="text-4xl" />
                    </label>
                </span>
            </div>
            <section class="content grid grid-cols-2 gap-10">
                <div class="editor flex flex-col gap-2">
                    <textarea v-model="posts[langPick].content" name="" id="" cols="30" rows="10"
                        class="content-input p-5 w-full border-2 border-dashed border-dark"></textarea>
                </div>
                <article class="content p-5 viewport w-full border-2 border-dashed border-dark"
                    v-html="md.render(posts[langPick].content)"></article>
            </section>
        </section>
        <button @click="createPost" class="btn"><a class="bg-white py-2 px-4 border-2 border-dark">Create post</a></button>
    </section>
</template>


<style>
.editor {
    cursor: default !important;
}
</style>