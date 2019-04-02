import React from 'react';
import './RegistrationForm.css'
// import LocalhostConn from '../api/LocalhostConn';

class RegistrationForm extends React.Component{
    state={name:"",address:"",addressLine2:"",postcode:"",outcode:"",type_of_food:"",rating:"",url:""}

    onSubmitClick =event=>{    //arraow function
        console.log("In Form Bar")
        console.log(this.state.Name)
        event.preventDefault();
        //console.log(this.state.term); //error watch udemy video no:84
        //send AJAX request to unsplash API
        this.props.onSubmit(this.state)
    }
   
    render(){
        console.log("In reg form render",this.state);
        return(
            <div className="register-form">
                <form className="ui form" onSubmit={this.onSubmitClick}>
                    <h3>RegistrationForm:</h3>
                    <div className="field">
                        <label> Name : </label> <input type="text" value={this.state.Name} onChange={(event)=>this.setState({name:event.target.value})}  />
                    </div>
                    <div className="field">
                        <label> Address :</label> <input type="text"  value={this.state.Address} onChange={(event)=>this.setState({address:event.target.value})} />
                    </div>
                    <div className="field">
                        <label> Address Line 2 : </label> <input type="text"  value={this.state.AddressLine2} onChange={(event)=>this.setState({addressLine2:event.target.value})} />
                    </div>
                    <div className="field">
                        <label> Postcode :</label> <input type="text"  value={this.state.Postcode} onChange={(event)=>this.setState({postcode:event.target.value})} />
                    </div>
                    <div className="field">
                        <label> Outcode :</label> <input type="text"  value={this.state.Outcode} onChange={(event)=>this.setState({outcode:event.target.value})} />
                    </div>
                    <div className="field">
                        <label> Type of Food : </label> <input type="text" value={this.state.TypeOfFood} onChange={(event)=>this.setState({type_of_food:event.target.value})} />
                    </div>
                    <div className="field">
                        <label> Rating :</label> <input type="text" value={this.state.Rating} onChange={(event)=>this.setState({rating:event.target.value})} />
                    </div>
                    <div className="field">
                        <label> URL : </label> <input type="text" value={this.state.URL} onChange={(event)=>this.setState({url:event.target.value})} />
                    </div>
                    <button className="ui button" type="submit" onClick={this.onSubmitClick}>Submit</button>
                </form>
            </div>
        );
    }
}

export default RegistrationForm;