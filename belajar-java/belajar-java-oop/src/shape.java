public class shape {

    int getCorner(){
        return 0;
    }

}

class Dua extends shape {
    int getCorner(){
        return 10;
    }

    int orangTuaCorner(){
        return super.getCorner();
    }

}

