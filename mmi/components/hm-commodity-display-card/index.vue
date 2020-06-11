<template>
  <div class="hm-commodity-display-card">
    <div class="itemWrap">
      <text class="title">{{ options.title }}</text>
      <image class="item" :src="options.item" />
      <div class="shanchulanmuWrap">
        <image class="shanchulanmu" :src="options.shanchulanmu" />
      </div>
      <div class="iconfriendsHelpWrap">
        <image class="iconfriendsHelp" :src="options.iconfriendsHelp" />
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'HmCommodityDisplayCard',
  props: {
    dataId: {
      type: String,
      default: 'hm-commodity-display-card'
    },
    options: {
      type: Object,
      default: function() {
        return {
          title: '凳子',
          item:
            './images/img_25366_0_0.png',
          shanchulanmu:
            './images/img_25366_0_2.png',
          iconfriendsHelp:
            './images/img_25366_0_1.png'
        };
      }
    }
  },
  data() {
    return {};
  },
  methods: {}
};
</script>
<style>
@import './index.response.css';
</style>
