<template>
  <div>
    <div>
      <div style="padding: 50px">
        <div v-show="!searchTerm" style="margin-bottom: 10px">
          categories
          <span v-for="cat in categories" :key="cat">
            <a-button
              style="margin-left: 5px"
              @click="getProductByCategory(cat)"
              class="name"
            >
              {{ cat }}
            </a-button>
          </span>
        </div>

        <a-row>
          <a-col
            :xl="{ span: 12 }"
            :lg="{ span: 12 }"
            :md="{ span: 12 }"
            :xs="{ span: 24 }"
          >
            <h2
              v-if="searchTerm"
              style="display: flex; justify-content: flex-start"
              class="product-title"
            >
              Search Results for: {{ searchTerm }}
            </h2>
            <h2
              v-if="!searchTerm"
              style="display: flex; justify-content: flex-start"
              class="product-title"
            >
              Проекты BI Group
            </h2>
          </a-col>
        </a-row>
        <a-input-search
          placeholder="input search text"
          enter-button
          style="width: 250px; margin-bottom: 1dvh"
          @search="onSearch"
          v-model="filter"
          @v-on:event="handle(event)"
        />
        <a-row :gutter="[20, 20]" v-if="items">
          <a-col
            class="gutter-row"
            v-for="item in items"
            :key="item.id"
            :xl="{ span: 6 }"
            :lg="{ span: 8 }"
            :md="{ span: 12 }"
            :xs="{ span: 24 }"
          >
            <a-card @click="openThisProduct(item)" hoverable>
              <img
                slot="cover"
                alt="example"
                :src="`https://realestate.enu.kz/api/static/${item.thumbnail[0]}`"
              />
              <a-card-meta>
                <span slot="title">
                  <a-tooltip>
                    <template slot="title">
                      {{ item.title }}
                    </template>
                    <span>{{ item.title.substring(0, 65) + "..." }}</span>
                  </a-tooltip>
                </span>
                <template slot="description">
                  <a-tooltip>
                    <template slot="title">
                      {{ item.description }}
                    </template>
                    <span>{{ item.description.substring(0, 65) + "..." }}</span>
                    <br />
                    <span class="price"> Цена: {{ item.price }} </span>
                  </a-tooltip>
                </template>
              </a-card-meta>
            </a-card>
          </a-col>
        </a-row>
      </div>
    </div>
    <!-- <div v-show="!items">
      <template>
        <a-row v-for="n in 4" :key="n" :gutter="[16, 32]">
          <a-col :span="6"> <a-skeleton /> </a-col>
          <a-col :span="6"> <a-skeleton /> </a-col>
          <a-col :span="6"> <a-skeleton /> </a-col>
          <a-col :span="6"> <a-skeleton /> </a-col>
        </a-row>
      </template>
    </div> -->
    <hr class="solid" style="margin-top: 30px; margin-bottom: 30px" />
    <iframe
      width="520"
      height="415"
      src="https://www.youtube.com/embed/616Vp9ml4Ac"
    >
    </iframe>
    <hr class="solid" style="margin-top: 30px; margin-bottom: 30px" />
    <div style="width: 100%; display: flex; flex-direction: row">
      <div
        style="
          align-items: center;
          align-content: center;
          align-self: center;
          width: 30%;
        "
      >
        <h1>Заказать звонок</h1>
        <h3 style="color: grey">Ответим на интересующие вас вопросы</h3>
        <h3 style="color: grey">и расскажем подробнее о проекте.</h3>
      </div>
      <div>
        <a-form style="margin-left: 30%; align-items: start">
          <a-form-item class="custom-label" label="">
            <div class="input-container">
              <span class="custom-label-text">Номер телефона</span>
              <a-input
                v-model="order.number"
                style="width: 400px; height: 45px"
                placeholder="+7 (___) ___ - __ - __"
              />
            </div>
          </a-form-item>

          <a-form-item class="custom-label" label="">
            <div class="input-container">
              <span class="custom-label-text">Город</span>
              <a-input
                v-model="order.city"
                style="width: 400px; height: 45px"
                placeholder="Астана"
              />
            </div>
          </a-form-item>
          <a-form-item class="custom-label" label="">
            <div class="input-container">
              <span class="custom-label-text">Ваше имя</span>
              <a-input
                v-model="order.name"
                style="width: 400px; height: 45px"
              />
            </div>
          </a-form-item>
          <a-form-item class="custom-label">
            <div class="input-container">
              <a-button
                type="primary"
                html-type="submit"
                style="width: 400px; height: 50px"
                @click="orderCall()"
                >Заказать звонок</a-button
              >
            </div>
          </a-form-item>
        </a-form>
      </div>
      <img
        height="348"
        width="351"
        src="https://s3.bi.group/biclick/content-manager/callback_f359377e8b.svg"
        alt=""
        class="c-gZsuB"
        style="margin-left: 20%"
      />
    </div>
  </div>
</template>

<script>
import { apiMixin } from "../mixins";
import axios from "axios";
import { mapState, mapMutations } from "vuex";

export default {
  created: async function () {
    // console.log('sssssssssssssssssssssss');
    if (this.$route.name === "Home") {
      await this.getAllProducts("Все");

      // return;
    }

    if (this.searchTerm) {
      const { data } = await axios(
        `https://dummyjson.com/products/search?q=${this.searchTerm}`
      );
      this.setItems(data.products);
    }
  },
  data() {
    return {
      filter: "",
      form: this.$form.createForm(this),
      order: {
        number: "",
        city: "",
        name: "",
      },
      // items: null,
      categories: ["Все", "Квартира", "Офис", "Участок", "Паркинг"],
    };
  },
  computed: {
    ...mapState(["items"]),
  },
  methods: {
    ...mapMutations([
      "setItems",
      "sortAscendingAction",
      "sortDescendingAction",
    ]),
    // setItemInStore(data) {},
    async orderCall() {
      console.log("order:", this.order);
      await axios.post("https://realestate.enu.kz/api/order", this.order, {
        withCredentials: true,
      });
      this.order = {};
    },
    async onSearch() {
      // console.log("filter:", this.filter);
      // console.log("event:", event);
      this.getAllProducts("Все", this.filter);
    },
    async getAllProducts(cat, text) {
      console.log("text:", text);
      if (text == "") text = undefined;
      const response = await axios(
        `https://realestate.enu.kz/api/residentialas/${cat}/${text}`,
        {
          withCredentials: true,
        }
      );
      console.log("data: ", response);

      // this.items = data;
      this.setItems(response.data.data);
    },
    async getProductByCategory(cat) {
      await this.getAllProducts(cat, "");
      // this.$router.push('/');
    },
    handle(event) {
      console.log("event:", event);
      this.getAllProducts("Все", event);
    },
    handleSubmit() {
      // Handle form submission here, for example, send data to the server
      console.log("Form submitted:", this.order);
      // Reset form fields
      this.form.resetFields();
      // You can also emit an event to notify parent component about the order
      this.$emit("order-placed", this.order);
    },

    openThisProduct(item) {
      this.$router.push({ name: "productDetails", params: { id: item._id } });
    },
  },
  mixins: [apiMixin],
  props: {
    searchTerm: String,
  },
};
</script>

<style scoped>
.input-container {
  display: flex;
  flex-direction: column; /* Align items in a column */
  align-items: flex-start;
}
.custom-label .ant-form-item-label {
  position: absolute;
  top: 0;
  left: 0;
}
.custom-label .custom-label-text {
  font-weight: bold; /* optional: make the label bold */
  margin-right: 8px; /* optional: add some space between label and input */
}
.name {
  font-weight: 600;
  text-transform: capitalize;
}
.ant-card-cover img {
  object-fit: contain;
  min-height: 280px;
  max-height: 280px;
  display: inline-block;
  padding: 12px;
}
.ant-card {
  height: 420px;
}
.price {
  color: orange;
}
.product-title {
  /* margin-bottom: 40px; */
  font-size: 35px;
  font-style: oblique;
}
</style>
