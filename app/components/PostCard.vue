<script setup>
import { ref, onBeforeMount } from 'vue'
import axios from 'axios'

const props = defineProps(['post'])


</script>

<template>
    <div class="wrapper max-w-[50rem] w-full mt-[-2px] px-10 border-y-2 border-dark border-dashed">
        <nuxt-link :to="`/blog/${post.id}`">
            <article class=" w-full border-x-2 border-dashed border-dark p-5 duration-300 hover:bg-dark hover:text-white relative pointer z-50">
                <h1 class="text-2xl">{{ post.title }}</h1>
                <p class=" line-clamp-3 text-sm">{{ post.description }}</p>
                <img :src="post.cover" alt="" class=" grayscale max-h-[10rem] w-full object-cover">
            </article>
        </nuxt-link>
    </div>
</template>


<style scoped>

article{
    /* cursor: default !important; */
}

</style>