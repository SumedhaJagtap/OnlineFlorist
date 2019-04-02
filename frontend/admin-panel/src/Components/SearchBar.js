import React from 'react';


//props only allows us to pass info from parent component to child
class SearchBar extends React.Component{
    // onInputChange(event){ //event is js object that contain info about event occurred
    //     console.log(event.target.value) //
    // }

    // onInputClick(){
    //     console.log("Input Clicked")
    // }

    onFormSubmit =event=>{    //arraow function
        //console.log("In search Bar")
        event.preventDefault();
        //console.log(this.state.term); //error watch udemy video no:84
        //send AJAX request to unsplash API
        this.props.onSubmit(this.state.term)
    }
    state = {term :''};
    render(){
        return (
            //() with this.onInputChange will call it automatically
        <div className="ui segment">
            <form className="ui form" onSubmit={this.onFormSubmit}>
                <div className="field">
                   <label> Restaurant Search</label>
                   <input 
                        type="text" 
                        value={this.state.term}
                        onChange={ (e) =>this.setState({term: e.target.value})}
                        
                        />  
                </div>
            </form>
        </div>
        );
    }
}

export default SearchBar;