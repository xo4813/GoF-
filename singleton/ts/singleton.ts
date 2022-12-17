class Singleton{
    private static instance: Singleton

    private constructor(){

    }

    public static getInstance(){
        return this.instance || (this.instance = new this())
    }

}