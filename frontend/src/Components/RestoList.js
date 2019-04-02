import React from 'react';
import faker from 'faker';
import LocalhostConn from '../api/LocalhostConn';
// import axios from 'axios';
import './RestoList.css';

class  RestoList extends React.Component{
    state={restos:[]};
    
    async componentDidMount(){
        const response =await LocalhostConn.get('',{
        });
        // console.log(response.data.responseData.resto)
        this.setState({restos:response.data.responseData.resto});
    };
    
    render(){
        const restos = this.state.restos.slice(0,10).map(({_id,name,postcode,address,rating,type_of_food})=>{
            // console.log({_id},{name})
            return <div key={_id}>
                Name : {name} <br />
                Address : {address} <br />
                Postcode {postcode} <br />
                Type Of Food : {type_of_food} <br />
                Rating : {rating} <br />
                Image : <img alt="food" src={faker.image.food()}></img>
                </div>
    
        });
        return <div className="resto-list"> 
        {restos}
         
     </div>
    }
    
};


export default RestoList;

