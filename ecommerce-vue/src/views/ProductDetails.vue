<template>
  <div>
    <ProductBanner :items="items"></ProductBanner>
    <div v-if="product">
      <a-space direction="vertical" size="large" style="padding: 50px 30px">
        <h2>{{ product.title }}</h2>
        <a-row :gutter="[24, 24]">
          <div class="align-item">
            <a-col :span="12">
              <div class="image-box">
                <img
                  :src="`http://localhost:2002/static/${product.thumbnail[0]}`"
                  alt=""
                  width="100%"
                />
              </div>
            </a-col>
            <a-col :span="12">
              <div class="product-detail">
                <h2>Price: {{ product.price }}</h2>
                <p>{{ product.description }}</p>
                <!-- <Counter :product="product"></Counter> -->
                <span>{{ product.builder }}</span>
              </div>
            </a-col>
          </div>
        </a-row>
      </a-space>
      <div class="image-box" style="width: 70%; margin: 0 auto">
        <GmapMap
          :center="center"
          :zoom="12"
          style="width: 100%; height: 400px"
          :options="{
            mapTypeId: 'terrain',
          }"
        >
        </GmapMap>
      </div>
      <div
        class="image-box"
        v-for="(schema, index) in product.schemas"
        :key="index"
      >
        <img
          :src="`http://localhost:2002/static/${schema}`"
          alt=""
          width="100%"
          style="margin-left: 10px; margin-right: 10px"
        />
      </div>
    </div>
    <div class="loader" v-else>
      <a-spin size="large" />
    </div>
  </div>
</template>

<script>
import ProductBanner from "@/components/ProductBanner.vue";
import axios from "axios";
import { apiMixin } from "../mixins";
import { gmapApi } from "vue2-google-maps";
export default {
  components: { ProductBanner },
  data() {
    return {
      product: null,
      center: { lat: 51.1694, lng: 71.4491 },
      items: [],
    };
  },
  computed: {},
  created: async function () {
    await gmapApi();
    const { id } = this.$route.params;

    if (!id) {
      this.$router.push(`/`);
      return;
    }

    const { data } = await axios(`http://localhost:2002/residential/${id}`);
    console.log({ data }.data.data);

    this.product = { data }.data.data;
    this.items = this.product.thumbnail;
    console.log("this.items:", this.items);
  },
  mixins: [apiMixin],
};
</script>

<style scoped>
.loader {
  height: calc(100vh - 64px);
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.image-box {
  border: 1px solid #cccccc;
  overflow: hidden;
  padding: 15px;
  border-radius: 5px;
}
.image-box img {
  max-width: 350px;
  object-fit: contain;
}
.product-detail {
  text-align: left;
}
.product-detail h2 {
  color: #747691;
  font-size: 30px;
}
.align-item {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
