*** Settings ***
Library     DatabaseLibrary

*** Test Cases ***
Orders rows
    Connect To Database     pymysql   toy     sealteam    sckshuhari    localhost     3306    
    # ${result}=     Query   Select * From orders
    # Log     ${result}
    Row Count Is Equal To X     Select * From orders    4