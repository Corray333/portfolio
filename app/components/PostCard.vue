<script setup>
import { ref, onBeforeMount } from 'vue'
import axios from 'axios'

const props = defineProps(['post'])


const reactions = ref(['mdi:thumb-up', 'mdi:thumb-down', 'mdi:heart', 'mdi:fire', 'fluent-emoji-high-contrast:moai', 'fluent-emoji-high-contrast:thinking-face'])


</script>

<template>
    <div class="wrapper max-w-[50rem] w-full mt-[-2px] px-10 border-y-2 border-dark border-dashed">
        <nuxt-link :to="`/blog/${post.id}`">
            <article class=" w-full border-x-2 border-dashed border-dark p-5 duration-300 hover:bg-dark hover:text-white relative pointer z-50">
                <h1 class="text-2xl">{{ post.title }}</h1>
                <p class=" line-clamp-3 text-sm">{{ post.description }}</p>
                <div class=" h-56 w-full relative">
                    <div class="flex gap-5 absolute left-0 bottom-0 z-50 text-white bg-dark text-xl p-2">
                        <p v-for="(reaction, i) of post.reactions" :key="i"><Icon :name="reactions[reaction.type]" /> {{ reaction.number }}</p>
                    </div>
                    <img :src="post.cover" alt="" class=" grayscale w-full h-full object-cover">
                </div>
            </article>
        </nuxt-link>
    </div>
</template>


<style scoped>

article{
    /* cursor: default !important; */
}

</style>