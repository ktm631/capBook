import React from 'react'
import { Button } from 'react-bootstrap'
import { Link } from "react-router-dom";

const Header = (state) => (
    <header className="App-header">
        <h1 className="App-title">
        <Link to={`/`} className="link" style={{ color: '#FFF' }}>
                    CapBook
                </Link>
        </h1>
        <div className="App-nav-buttons">
            <Button variant="light" className="">                
                <Link to={`/login`} style={{ color: '#FFF' }}>
                    Sign In
                </Link></Button>
            <Button variant="light">
            <Link to={`/register`} style={{ color: '#FFF' }}>
            
            Sign up
            </Link>
            </Button>
        </div>
        <div className="clear"></div>
    </header>
);

export default Header;
