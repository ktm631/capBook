import React, { Component } from 'react';
import './App.css';
import Login from "./components/Login";
import Register from "./components/Register";
import BookTable from "./components/BookTable";
import Header from "./components/Header";
import { BrowserRouter, Route } from 'react-router-dom';
import BookElement from './components/BookElement';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isAuth: false
        }
    }
    render() {
        return (
            <BrowserRouter>
                <div className="App">
                    <Header/>
                    <Route exact path="/" component={BookTable}/>
                    <Route path="/login" component={Login} />
                    <Route path="/register" component={Register} />
                    <Route path="/book/:id" component={BookElement} />
                </div>
            </BrowserRouter>
        );
    }
}

export default App;
