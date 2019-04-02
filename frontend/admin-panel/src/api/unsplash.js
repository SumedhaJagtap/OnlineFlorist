import axios from 'axios';


// const searchImages = (term) =>{

// }

export default axios.create({
    baseURL:"https://api.unsplash.com", 
    headers:{
        Authorization:'Client-ID e08634caaebb14966e39fd788833734d31021527308a3f6a0a94a7c8aae33edf'
    }
});