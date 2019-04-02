import axios from 'axios';


// const searchImages = (term) =>{

// }

export default axios.create({
    baseURL:"http://localhost:8080/restaurantservice/restaurant/"
});