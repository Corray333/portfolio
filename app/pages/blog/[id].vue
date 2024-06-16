<script setup>
import { ref, onBeforeMount } from 'vue'
import axios from 'axios'
import markdownit from 'markdown-it'

const md = markdownit({
    html: true,
    linkify: true,
    typographer: true,
    xhtmlOut: true,
})

const post = ref(null)
const route = useRoute()

const loadPost = async () => {
    try {
        let { data } = await axios.get(`/api/posts/${route.params.id}?lang=${localStorage.getItem('lang')}`)
        post.value = data.post
        loadReactions()
    } catch (err) {
        console.log(err)
    }
}

const currentReactions = ref([])
const reactionSet = ref(-1)
const loadReactions = () => {
    currentReactions.value = localStorage.getItem('reactions') ? JSON.parse(localStorage.getItem('reactions')) : {}
    for (let i = 0; i < currentReactions.value.length; i++) {
        if (currentReactions.value[i].id == post.value.id) {
            reactionSet.value = currentReactions.value[i].type
            break
        }
    }
}

const setReaction = (type) => {
    if (reactionSet.value == type) {
        return
    }
    if (reactionSet.value == -1) {
        currentReactions.value.push({ id: post.value.id, type })
        localStorage.setItem('reactions', JSON.stringify(currentReactions.value))
        reactionSet.value = type

        post.value.reactions[type].number++
        try {
            axios.post(`/api/posts/${post.value.id}/reactions`, {
                increment: type
            })
        } catch (err) {
            console.log(err)
        }
    } else {
        for (let i = 0; i < currentReactions.value.length; i++) {
            if (currentReactions.value[i].id == post.value.id) {
                currentReactions.value[i].type = type
                // TODO: send to server
                let old = reactionSet.value
                if (post.value.reactions[reactionSet.value].number > 0) post.value.reactions[reactionSet.value].number--
                reactionSet.value = type

                post.value.reactions[type].number++
                try {
                    axios.post(`/api/posts/${post.value.id}/reactions`, {
                        increment: type,
                        decrement: old
                    })
                } catch (err) {
                    console.log(err)
                }

                localStorage.setItem('reactions', JSON.stringify(currentReactions.value))
                break
            }
        }
    }


}

onBeforeMount(() => {
    loadPost()
})

onMounted(() => {


    document.querySelector('.cursor').style.transform = 'translateX(-50%) translateY(-50%) scale(1)'

})

const reactions = ref(['mdi:thumb-up', 'mdi:thumb-down', 'mdi:heart', 'mdi:fire', 'fluent-emoji-high-contrast:moai', 'fluent-emoji-high-contrast:thinking-face'])

</script>

<template>
    <section v-if="post" class=" w-full min-h-screen p-20 flex flex-col items-center ">
        <article class="w-full max-w-[60rem]">
            <article class="header relative">
                <nuxt-link to="/blog" class="p-2 absolute pointer rounded-full right-full mr-5 hover:bg-dark group duration-300"><Icon class="text-dark duration-300 group-hover:text-white text-4xl" name="mdi:undo" /></nuxt-link>
                <img :src="post.cover" class="w-full object-cover mb-5" alt="">
                <div class="flex gap-5 my-2">
                    <p v-for="(tag, i) of post.tags" :key="i" class=" text-white bg-dark p-2 text-sm">{{ tag }}</p>
                </div>
                <h1 class="text-4xl">{{ post.title }}</h1>
                <p>{{ post.description }}</p>
            </article>
            <hr class="my-5">
            <article v-html="md.render(post.content)" class="content"></article>
            <div class="flex gap-5 left-0 bottom-0 z-50 text-xl p-2 relative">
                <p v-for="(reaction, i) of post.reactions" :key="i" @click="setReaction(reaction.type)"
                    class="pointer hover:bg-dark hover:text-white duration-300 p-2 rounded-full"
                    :class="reaction.type == reactionSet ? 'bg-dark text-white' : ''">
                    <Icon :name="reactions[reaction.type]" /> {{ reaction.number }}
                </p>
            </div>
        </article>
    </section>
</template>

<style></style>