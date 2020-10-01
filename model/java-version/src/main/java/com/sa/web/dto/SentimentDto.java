package com.sa.web.dto;

public class SentimentDto {

    private String sentence;
    private float polarity;
    private String version;

    public SentimentDto() {
        this.version = "java";
    }

    public SentimentDto(String sentence, float polarity) {
        this.sentence = sentence;
        this.polarity = polarity;
        this.version = "java";
    }

    public String getSentence() {

        return sentence;
    }

    public void setSentence(String sentence) {
        this.sentence = sentence;
    }

    public String getVersion() {

        return version;
    }

    public float getPolarity() {
        return polarity;
    }

    public void setPolarity(float polarity) {
        this.polarity = polarity;
    }

    @Override
    public String toString() {
        return "SentimentDto{" +
                "sentence='" + sentence + '\'' +
                ", polarity='" + polarity + '\'' +
                ", version=" + version +  
                '}';
    }
}
