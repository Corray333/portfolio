<script setup>
import {ref, onBeforeMount} from 'vue'
import axios from 'axios'
import {useRouter} from 'vue-router'

const router = useRouter()

const posts = ref([])

const loadPosts = async ()=>{
    try{
        // TODO: add lang
        let lang = localStorage.getItem('lang')
        let {data} = await axios.get(`/api/posts?lang=${lang}`)
        posts.value = data.posts
    } catch(err){
        console.log(err)
    }
}

const refreshTokens = async ()=>{
    try{
        let {data} = await axios.get(`/api/admin/refresh`,{
            headers:{
                Rerfresh:localStorage.getItem('Refresh')
            }
        })
        localStorage.setItem('Authorization', data.authorization)
        localStorage.setItem('Refresh', data.refresh)
    }catch(err){
        console.log(err)
    }
}

onBeforeMount(async ()=>{
    if (!localStorage.getItem('Authorization')){
        router.push(`admin/login`)
        return
    }
    await refreshTokens()
    loadPosts()
})

</script>

<template>
  <section class="p-20 w-full min-h-screen">
    <h1 class=" text-[4rem] text-center flex items-center gap-2 justify-center">Admin <button @click="router.push(`/admin/new-post`)" class="z-50 relative pointer text-4xl  duration-300 hover:bg-dark hover:text-white p-2 rounded-full w-12 h-12 flex justify-center items-center"><Icon name="mdi:plus" /></button></h1>
    <div class="posts flex flex-col items-center">
        <div class="wrapper  w-full max-w-[50rem] px-10">
            <div class="top h-10 border-x-2 border-dashed border-dark"></div>
        </div>
        <PostCard v-for="(post,i) of posts" :key="i" :post="post" :link="`/admin/${post.id}`" />
        <div class="wrapper  w-full max-w-[50rem] px-10">
            <div class="top h-10 border-x-2 border-dashed border-dark"></div>
        </div>
    </div>
  </section>
</template>


<style>

</style>