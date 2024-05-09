<script setup>
import { ref, onBeforeMount } from 'vue'
import axios from 'axios'
import markdownit from 'markdown-it'

const md = markdownit({
    html: true,
    linkify: true,
    typographer: true,
    xhtmlOut:     true,
})

const post = ref(null)
const route = useRoute()

const loadPost = async () => {
    try {
        let { data } = await axios.get(`/api/posts/${route.params.id}?lang=${localStorage.getItem('lang')}`)
        post.value = data.post
    } catch (err) {
        console.log(err)
    }
}

onBeforeMount(() => {
    loadPost()
})


</script>

<template>
    <section v-if="post" class=" w-full min-h-screen p-20 flex flex-col items-center ">
        <article class="w-full max-w-[60rem]">
            <article class="header">
                <img :src="post.cover" class="w-full h-48 object-cover" alt="">
                <h1 class="text-2xl">{{ post.title }}</h1>
                <p>{{ post.description }}</p>
            </article>
            <hr>
            <article v-html="md.render(post.content)" class="content"></article>
        </article>
    </section>
</template>

<style>

</style>