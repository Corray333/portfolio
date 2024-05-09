<script setup>
import { ref, onMounted } from 'vue'

const cursor = ref(null)

onMounted(() => {
  if (!localStorage.getItem('lang')) localStorage.setItem('lang','eng')

  document.addEventListener('mouseover', (e) => {
    if (e.target.closest('.pointer')) {
      document.querySelector('.cursor').style.transform = 'translateX(-50%) translateY(-50%) scale(0)'
    }
  })
  document.addEventListener('mouseout', (e) => {
    if (e.target.closest('.pointer')) {
      document.querySelector('.cursor').style.transform = 'translateX(-50%) translateY(-50%) scale(1)'
    }
  })
})

onMounted(() => {
  window.addEventListener('mousemove', (e) => {
    cursor.value.style.top = window.scrollY + e.clientY + 'px'
    cursor.value.style.left = window.scrollX + e.clientX + 'px'
  })
})


let slideUp = (target, duration = 500) => {
  target.style.transitionProperty = 'height, margin, padding';
  target.style.transitionDuration = duration + 'ms';
  target.style.boxSizing = 'border-box';
  target.style.height = target.offsetHeight + 'px';
  target.offsetHeight;
  target.style.overflow = 'hidden';
  target.style.height = 0;
  target.style.paddingTop = 0;
  target.style.paddingBottom = 0;
  target.style.marginTop = 0;
  target.style.marginBottom = 0;
  window.setTimeout(() => {
    target.style.display = 'none';
    target.style.removeProperty('height');
    target.style.removeProperty('padding-top');
    target.style.removeProperty('padding-bottom');
    target.style.removeProperty('margin-top');
    target.style.removeProperty('margin-bottom');
    target.style.removeProperty('overflow');
    target.style.removeProperty('transition-duration');
    target.style.removeProperty('transition-property');
    //alert("!");
    headerOpened.value = false
  }, duration);
}

/* SLIDE DOWN */
let slideDown = (target, duration = 500) => {

  target.style.removeProperty('display');
  let display = window.getComputedStyle(target).display;
  if (display === 'none') display = 'flex';
  target.style.display = display;
  let height = target.offsetHeight;
  target.style.overflow = 'hidden';
  target.style.height = 0;
  target.style.paddingTop = 0;
  target.style.paddingBottom = 0;
  target.style.marginTop = 0;
  target.style.marginBottom = 0;
  target.offsetHeight;
  target.style.boxSizing = 'border-box';
  target.style.transitionProperty = "height, margin, padding";
  target.style.transitionDuration = duration + 'ms';
  target.style.height = height + 'px';
  target.style.removeProperty('padding-top');
  target.style.removeProperty('padding-bottom');
  target.style.removeProperty('margin-top');
  target.style.removeProperty('margin-bottom');
  window.setTimeout(() => {
    target.style.removeProperty('height');
    target.style.removeProperty('overflow');
    target.style.removeProperty('transition-duration');
    target.style.removeProperty('transition-property');
  }, duration);
}
var slideToggle = ( duration = 500) => {
  let target = document.querySelector('.navbar'); 
  if (window.getComputedStyle(target).display === 'none') {
    headerOpened.value = true
    slideDown(target, duration);
  } else {
    slideUp(target, duration);
  }
  target = document.querySelector('.nav_bg'); 
  if (window.getComputedStyle(target).display === 'none') {
    slideDown(target, duration);
  } else {
    slideUp(target, duration);
  }
}


const headerOpened = ref(false)

</script>
<template>
  <section class=" overflow-hidden cursor-none relative z-10">
    <span ref="cursor"
      class="cursor z-40 hidden sm:flex transition-transform duration-200 absolute bg-white mix-blend-difference w-12 h-12 justify-center items-center rounded-full">
      <Icon class="text-xl text-black" name="fluent:cursor-16-filled" />
    </span>
    <button @click="slideToggle()" class="burger pointer fixed right-0 top-0 mt-10 mr-10 p-2 duration-300 group hover:bg-white rounded-md">
        <span class="group-hover:bg-dark"></span><span class="group-hover:bg-dark"></span><span class="group-hover:bg-dark"></span>
    </button>
    <navbar class=" navbar fixed hidden" @route="slideToggle()"/>
    <div class="nav_bg fixed bg-dark w-screen h-screen z-30 hidden">
      <div class="marquee text-white uppercase text-[4rem] whitespace-nowrap absolute right-0 top-0 rotate-45 font-bold" style="line-height: 0.8; width: 180vh;">
        <p class="marquee1">navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation</p>
        <p class="marquee2">navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation navigation</p>
      </div>
    </div>
    <NuxtPage :style="headerOpened ? 'z-index:-1; position:relative':''" />
  </section>
</template>


<style setup>

.marquee1 {
  animation: marquee1 30s linear infinite;
}
.marquee2 {
  animation: marquee2 30s linear infinite;
}


@keyframes marquee1 {
  0% { transform: translateX(0); }
  50% { transform: translateX(-100%); }
  100% { transform: translateX(0); }
}

@keyframes marquee2 {
  0% { transform: translateX(-100%); }
  50% { transform: translateX(0); }
  100% { transform: translateX(-100%); }
}
.pointer{
  cursor: none;
}

.burger{
        @apply flex flex-col gap-1;
        z-index: 100;
        mix-blend-mode: difference;
        span{
            @apply block w-6 h-1 bg-white rounded-full;
        }
    }

section {
  background-color: #fff;
}

.cursor {
  transform: translateX(-50%) translateY(-50%);
  cursor: none;
}
</style>