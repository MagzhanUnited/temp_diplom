<script setup>
import { FilterMatchMode } from 'primevue/api';
import { ref, onMounted } from 'vue';
import { ProductService } from '@/service/ProductService';
import { useToast } from 'primevue/usetoast';
import ResidentService from '../../service/ResidentService';

const toast = useToast();

const products = ref(null);
const productDialog = ref(false);
const deleteProductDialog = ref(false);
const deleteProductsDialog = ref(false);
const product = ref({});
const selectedProducts = ref(null);
const dt = ref(null);
const filters = ref({});
const submitted = ref(false);
const fileUpload = ref(null);
const schemasUpload = ref(null);
const statuses = ref([
    { label: 'INSTOCK', value: 'instock' },
    { label: 'LOWSTOCK', value: 'lowstock' },
    { label: 'OUTOFSTOCK', value: 'outofstock' }
]);
const customBase64Uploader = async (event) => {
    console.log('ERE');
    const file = event.files[0];
    const reader = new FileReader();
    let blob = await fetch(file.objectURL).then((r) => r.blob()); //blob:url

    reader.readAsDataURL(blob);

    reader.onloadend = function () {
        const base64data = reader.result;
        console.log('base64data:', base64data);
    };
};
const productService = new ProductService();

const getBadgeSeverity = (inventoryStatus) => {
    switch (inventoryStatus.toLowerCase()) {
        case 'instock':
            return 'success';
        case 'lowstock':
            return 'warning';
        case 'outofstock':
            return 'danger';
        default:
            return 'info';
    }
};
const init = async () => {
    const response = await ResidentService.getResidents('Все');

    products.value = response.data.data;
    console.log(products.value);
};
onMounted(async () => {
    // productService.getProducts().then((data) => (products.value = data));
    await init();
});
const formatCurrency = (value) => {
    return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
};

const openNew = () => {
    product.value = {};
    submitted.value = false;
    productDialog.value = true;
};

const hideDialog = () => {
    productDialog.value = false;
    submitted.value = false;
};

const onUpload = (event) => {
    // Handle the file upload response
    console.log('File uploaded', event);
};
const saveProduct = async () => {
    console.log(product.value);
    submitted.value = true;
    if (product.value.title && product.value.title.trim() && product.value.price) {
        if (product.value.id) {
            product.value.inventoryStatus = product.value.inventoryStatus.value ? product.value.inventoryStatus.value : product.value.inventoryStatus;
            products.value[findIndexById(product.value.id)] = product.value;
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Updated', life: 3000 });
        } else {
            // product.value.id = createId();
            // product.value.code = createId();
            // product.value.image = 'product-placeholder.svg';
            // product.value.inventoryStatus = product.value.inventoryStatus ? product.value.inventoryStatus.value : 'INSTOCK';
            // products.value.push(product.value);
            const uploadUrl = `/upload`;

            const fileupload = fileUpload.value;
            const schemasupload = schemasUpload.value;
            if (fileupload && fileupload.$options) {
                fileupload.$options.url = uploadUrl;

                // Trigger the file upload
                // console.log('fileUpload1:', fileupload);
                // fileupload.upload();
                // console.log('fileUpload2:', fileupload);
                console.log('fileupload.files:', fileupload.files);
                product.value.thumbnail = [];
                for (const file of fileupload.files) {
                    const formData = new FormData();
                    product.value.thumbnail.push(file.name);
                    console.log('file.name:', file.name);
                    formData.append('demo[]', file);
                    await ResidentService.upload(formData);
                }
                //  = fileupload.files[0].name;
            } else {
                console.error('FileUpload or options is undefined. Check your component and template.');
            }
            if (schemasupload && schemasupload.$options) {
                schemasupload.$options.url = uploadUrl;

                // Trigger the file upload
                // console.log('schemasupload1:', schemasupload);
                // schemasupload.upload();
                // console.log('schemasupload2:', schemasupload);
                console.log('schemasupload.files:', schemasupload.files);
                product.value.schemas = [];
                for (const file of schemasupload.files) {
                    const formData = new FormData();
                    product.value.schemas.push(file.name);
                    console.log('file.name:', file.name);
                    formData.append('demo[]', file);
                    await ResidentService.upload(formData);
                }
                //  = schemasupload.files[0].name;
            } else {
                console.error('FileUpload or options is undefined. Check your component and template.');
            }
            console.log(product.value);
            await ResidentService.postResident(product.value);
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Created', life: 3000 });
            await init();
        }
        productDialog.value = false;
        product.value = {};
    }
};

const editProduct = (editProduct) => {
    product.value = { ...editProduct };
    productDialog.value = true;
};

const confirmDeleteProduct = (editProduct) => {
    product.value = editProduct;
    deleteProductDialog.value = true;
};

const deleteProduct = async () => {
    // products.value = products.value.filter((val) => val.id !== product.value.id);
    console.log('product.value._id:', product.value._id);
    await ResidentService.deleteResident(product.value._id);
    deleteProductDialog.value = false;
    product.value = {};
    await init();
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 });
};

const findIndexById = (id) => {
    let index = -1;
    for (let i = 0; i < products.value.length; i++) {
        if (products.value[i].id === id) {
            index = i;
            break;
        }
    }
    return index;
};

const createId = () => {
    let id = '';
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    for (let i = 0; i < 5; i++) {
        id += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return id;
};

const exportCSV = () => {
    dt.value.exportCSV();
};

const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};
const deleteSelectedProducts = () => {
    products.value = products.value.filter((val) => !selectedProducts.value.includes(val));
    deleteProductsDialog.value = false;
    selectedProducts.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 });
};

const initFilters = () => {
    filters.value = {
        global: { value: null, matchMode: FilterMatchMode.CONTAINS }
    };
};
initFilters();
</script>

<template>
    <div class="grid">
        <div class="col-12">
            <div class="card">
                <Toolbar class="mb-4">
                    <template v-slot:start>
                        <div class="my-2">
                            <Button label="New" icon="pi pi-plus" class="mr-2" severity="success" @click="openNew" />
                            <!-- <Button label="Delete" icon="pi pi-trash" severity="danger" @click="confirmDeleteSelected" :disabled="!selectedProducts || !selectedProducts.length" /> -->
                        </div>
                    </template>

                    <template v-slot:end>
                        <FileUpload mode="basic" accept="image/*" :maxFileSize="1000000" label="Import" chooseLabel="Import" class="mr-2 inline-block" />
                        <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)" />
                    </template>
                </Toolbar>

                <DataTable
                    ref="dt"
                    :value="products"
                    v-model:selection="selectedProducts"
                    dataKey="id"
                    :paginator="true"
                    :rows="10"
                    :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[5, 10, 25]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products"
                >
                    <template #header>
                        <div class="flex flex-column md:flex-row md:justify-content-between md:align-items-center">
                            <h5 class="m-0">Manage Products</h5>
                            <span class="block mt-2 md:mt-0 p-input-icon-left">
                                <i class="pi pi-search" />
                                <InputText class="w-full sm:w-auto" v-model="filters['global'].value" placeholder="Search..." />
                            </span>
                        </div>
                    </template>

                    <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
                    <Column field="title" header="title" :sortable="true" headerStyle="width:14%; min-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Code</span>
                            {{ slotProps.data.title }}
                        </template>
                    </Column>
                    <Column field="category" header="category" :sortable="true" headerStyle="width:14%; min-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Name</span>
                            {{ slotProps.data.category }}
                        </template>
                    </Column>
                    <Column header="Image" headerStyle="width:14%; min-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Image</span>
                            <img :src="'http://194.4.58.178:8000/static/' + slotProps.data.thumbnail[0]" :alt="slotProps.data.thumbnail[0]" class="shadow-2" width="100" />
                        </template>
                    </Column>
                    <Column field="discountPercentage" header="Скидка" :sortable="true" headerStyle="width:14%; min-width:8rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">discountPercentage</span>
                            {{ slotProps.data.discountPercentage }}%
                        </template>
                    </Column>
                    <Column field="price" header="price" :sortable="true" headerStyle="width:14%; min-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Price</span>
                            {{ slotProps.data.price }}
                        </template>
                    </Column>
                    <Column field="rating" header="Reviews" :sortable="true" headerStyle="width:14%; min-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Rating</span>
                            <Rating :modelValue="slotProps.data.rating" :readonly="true" :cancel="false" />
                        </template>
                    </Column>
                    <Column field="builder" header="Builder" :sortable="true" headerStyle="width:14%; min-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Builder</span>
                            {{ slotProps.data.builder }}
                        </template>
                    </Column>
                    <Column headerStyle="min-width:10rem;">
                        <template #body="slotProps">
                            <Button icon="pi pi-pencil" class="mr-2" severity="success" rounded @click="editProduct(slotProps.data)" />
                            <Button icon="pi pi-trash" class="mt-2" severity="warning" rounded @click="confirmDeleteProduct(slotProps.data)" />
                        </template>
                    </Column>
                </DataTable>

                <Dialog v-model:visible="productDialog" :style="{ width: '450px' }" header="Product Details" :modal="true" class="p-fluid">
                    <img :src="'http://194.4.58.178:8000/static/' + product.thumbnail[0]" :alt="product.thumbnail[0]" v-if="product.thumbnail" width="150" class="mt-0 mx-auto mb-5 block shadow-2" />
                    <!-- <FileUpload v-else mode="basic" name="demo[]" url="http://194.4.58.178:8000/upload" accept="image/*" :maxFileSize="1000000" customUpload @uploader="customBase64Uploader" style="margin-top: 2dvh; margin-bottom: 2dvh" /> -->
                    <FileUpload v-else mode="advanced" name="demo[]" url="http://194.4.58.178:8000/upload" :showUploadButton="false" :multiple="true" ref="fileUpload" accept="image/*" :maxFileSize="10000000" @uploader="customBase64Uploader">
                        <template #empty>
                            <p>Перетащите картинки дома сюда.</p>
                        </template>
                    </FileUpload>

                    <div class="field">
                        <label for="name">Title</label>
                        <InputText id="name" v-model.trim="product.title" required="true" autofocus :class="{ 'p-invalid': submitted && !product.title }" />
                        <small class="p-invalid" v-if="submitted && !product.title">Name is required.</small>
                    </div>
                    <div class="field">
                        <label for="description">Description</label>
                        <Textarea id="description" v-model="product.description" required="true" rows="3" cols="20" />
                    </div>

                    <div class="formgrid grid">
                        <div class="field col">
                            <label for="discountPercentage">Скидка</label>
                            <InputNumber id="discountPercentage" v-model="product.discountPercentage" :class="{ 'p-invalid': submitted && !product.discountPercentage }" :required="true" />
                            <small class="p-invalid" v-if="submitted && !product.discountPercentage">Скидка is required.</small>
                        </div>
                        <div class="field col">
                            <label for="rating">Rating</label>
                            <InputNumber id="rating" v-model="product.rating" />
                        </div>
                    </div>

                    <div class="field">
                        <label class="mb-3">Category</label>
                        <div class="formgrid grid">
                            <div class="field-radiobutton col-6">
                                <RadioButton id="category1" name="category" value="Квартира" v-model="product.category" />
                                <label for="category1">Квартира</label>
                            </div>
                            <div class="field-radiobutton col-6">
                                <RadioButton id="category2" name="category" value="Участок" v-model="product.category" />
                                <label for="category2">Участок</label>
                            </div>
                            <div class="field-radiobutton col-6">
                                <RadioButton id="category3" name="category" value="Офис" v-model="product.category" />
                                <label for="category3">Офис</label>
                            </div>
                            <div class="field-radiobutton col-6">
                                <RadioButton id="category4" name="category" value="Паркинг" v-model="product.category" />
                                <label for="category4">Паркинг</label>
                            </div>
                        </div>
                    </div>

                    <div class="formgrid grid">
                        <div class="field col">
                            <label for="price">Price</label>
                            <InputText id="price" v-model="product.price" :class="{ 'p-invalid': submitted && !product.price }" :required="true" />
                            <small class="p-invalid" v-if="submitted && !product.price">Price is required.</small>
                        </div>
                        <div class="field col">
                            <label for="builder">Builder</label>
                            <InputText id="builder" v-model="product.builder" />
                        </div>
                    </div>
                    <!-- <FileUpload v-if="!product.thumbnail" mode="advanced" name="demo[]" url="http://194.4.58.178:8000/upload" :showUploadButton="false" :multiple="true" ref="fileUpload" accept="image/*" :maxFileSize="10000000" @upload="onUpload">
                        <template #empty>
                            <p>Перетащите схему дома сюда.</p>
                        </template>
                    </FileUpload> -->
                    <FileUpload
                        v-if="!product.thumbnail"
                        mode="advanced"
                        name="demo[]"
                        url="http://194.4.58.178:8000/upload"
                        :showUploadButton="false"
                        :multiple="true"
                        ref="schemasUpload"
                        accept="image/*"
                        :maxFileSize="10000000"
                        @uploader="customBase64Uploader"
                    >
                        <template #empty>
                            <p>Перетащите картинки дома сюда.</p>
                        </template>
                    </FileUpload>
                    <template #footer>
                        <Button label="Cancel" icon="pi pi-times" text="" @click="hideDialog" />
                        <Button label="Save" icon="pi pi-check" text="" @click="saveProduct" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteProductDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="product"
                            >Are you sure you want to delete <b>{{ product.title }}</b
                            >?</span
                        >
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" text @click="deleteProductDialog = false" />
                        <Button label="Yes" icon="pi pi-check" text @click="deleteProduct" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteProductsDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="product">Are you sure you want to delete the selected products?</span>
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" text @click="deleteProductsDialog = false" />
                        <Button label="Yes" icon="pi pi-check" text @click="deleteSelectedProducts" />
                    </template>
                </Dialog>
            </div>
        </div>
    </div>
</template>
