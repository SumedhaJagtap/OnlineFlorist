import React from 'react';
import faker from 'faker';
import './SearchedResto.css';
// import Unsplash from '../api/unsplash';


const SearchedResto = (props) =>{
//     console.log(props.restos)
//     restoImage=faker.image.food()
//     console.log(restoImage)
//     // const response =await Unsplash.get('/search/photos',{
//     //         params:{
//     //             query:"restaurant"
//     //         }
//     //     });
//     // console.log(response)
//     // console.log("hey",response.data.responseData)
//     // restoImage=response.data.urls.small
//    console.log("Meeeeeeeeeee")
    const restos =props.restos.map(({id,name,postcode,address,rating,type_of_food}) => {
        return <div key={id} className="searched-resto">
            Name : {name} <br />
            Address : {address} <br />
            Postcode {postcode} <br />
            Type Of Food : {type_of_food} <br />
            Rating : {rating} <br />
            Image : <img alt="food" src={faker.image.food()}></img>
         
        </div>
    })
    // console.log("Hey",restoImage)
    return <div className="resto-list"> 
       {restos}
        
    </div>
};


export default SearchedResto;

