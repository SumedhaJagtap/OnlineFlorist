import React from 'react';
import ReactDOM from 'react-dom';
// import faker from 'faker';
// import unsplash from './api/unsplash';
import SearchBar from './Components/SearchBar';
import RestoList from './Components/RestoList';
import LocalhostConn from './api/LocalhostConn';
// import Spinner from './Components/LoaderSpinner';
import SearchedResto from './Components/SearchedResto';
import './Components/errors.css';
import RegistrationForm from './Components/registrationForm';
// import './Components/RegistrationForm.css'
//import axios from 'axios';
class RestoApp extends React.Component{

    state={restos:[],err:""};
    // onSearchSubmit(term){  //promise based syntax
    //     //console.log(term)
    //     axios.get('https://api.unsplash.com/search/photos',{
    //         params:{
    //             query:term
    //         },
    //         headers:{
    //             Authorization:'Client-ID e08634caaebb14966e39fd788833734d31021527308a3f6a0a94a7c8aae33edf'
    //         }
    //     }).then((response) =>{
    //         console.log(response.data.results);
    //     }
    //     );
    // }



    //below will give error for this that why following arraw function is used
    // async onSearchSubmit(term){
    //     //console.log(term)
    //     const response =await axios.get('https://api.unsplash.com/search/photos',{
    //         params:{
    //             query:term
    //         },
    //         headers:{
    //             Authorization:'Client-ID e08634caaebb14966e39fd788833734d31021527308a3f6a0a94a7c8aae33edf'
    //         }
    //     });
    //     console.log(this);
    //     this.setState({images:response.data.results});
    // }
    onRegisterSubmit= async(event) =>{
        console.log(event)
        const response =await LocalhostConn.post('',{
                name:event.name,
                address:event.address,
                addressLine2:event.addressLine2,
                url:event.url,
                outcode:event.outcode,
                postcode:event.postcode,
                rating:event.rating,
                type_of_food:event.type_of_food
        });
        console.log(response.data.responseData)
        // if(response.data.responseData.resto===null){
        //    // if(response.data.responseData.resto.length===0){
        //         this.setState({err:"No Restaurant Found"})
        //     //}
        // }
        // this.setState({restos:response.data.responseData.resto});
    };
     onSearchSubmit = async (term) =>{
        const response =await LocalhostConn.get('',{
            params:{
                name:term
            }
        });
        // console.log(response.data.responseData)
        if(response.data.responseData.resto===null){
           // if(response.data.responseData.resto.length===0){
                this.setState({err:"No Restaurant Found"})
            //}
        }
        this.setState({restos:response.data.responseData.resto});
    }
    // onSearchSubmit(term){
    //     axios.get('http://localhost:8080/restaurantservice/restaurant/',
    //         {
    //         params:{
    //             name:term
    //         }
    //     }).then((response) =>{
    //                 console.log(response.data.responseData);
    //             }
    //             );
    //         }

    // componentDidMount(){
    //   var loadpage = async (term) =>{
    //         const response =await LosthostConn.get('',{
    //         });
    //         console.log(response.data.responseData)
    //         this.setState({restos:response.data.responseData.resto});
    //     };
    //   return <div style="display:grid" >
    //   {loadpage()}
    //   </div>
    // }
    renderContent(){
        if (this.state.restos){
            if (this.state.restos.length!==0){
                return(
                 <div>
    
                        <h3> {this.state.restos.length} Searches Found </h3>
                        <SearchedResto restos={this.state.restos} />
                     </div>   
                );
            }
        }else if(this.state.err!==""){
                return(
                    <div className="not-found">
                        <h4> {this.state.err} </h4>
                        Please Enter Valid Reastaurant Name
                    </div>   
                   )
            }
        
       
        

    }

    render(){
       return( 
                <div className="ui container" style={{marginTop : '10px'}}> 
                    <SearchBar onSubmit={this.onSearchSubmit} />
                    
                {this.renderContent()}
                 <div>  <RestoList />  </div>
                 <RegistrationForm onSubmit={this.onRegisterSubmit} />

           </div>
       );
    }
    
};

ReactDOM.render(
    <RestoApp />,
    document.querySelector('#root')
);


