<script setup>
import {ref, onBeforeMount} from 'vue'
import axios from 'axios'

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

onBeforeMount(()=>{
    loadPosts()
})

</script>

<template>
  <section class="p-20 w-full min-h-screen">
    <h1 class=" text-[4rem] text-center">Blog</h1>
    <div class="posts flex flex-col items-center">
        <div class="wrapper  w-full max-w-[50rem] px-10">
            <div class="top h-10 border-x-2 border-dashed border-dark"></div>
        </div>
        <PostCard v-for="(post,i) of posts" :key="i" :post="post" :link="`/blog/${post.id}`" />
        <div class="wrapper  w-full max-w-[50rem] px-10">
            <div class="top h-10 border-x-2 border-dashed border-dark"></div>
        </div>
    </div>
  </section>
</template>


<style>

</style>