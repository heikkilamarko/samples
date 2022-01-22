using System.Globalization;
using System.Text.Encodings.Web;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Text.Unicode;
using CsvHelper;
using CsvHelper.Configuration;
using CsvHelper.Configuration.Attributes;

namespace csv2json;

class Processor
{
    public void Process(string inFilePath, string outFilePath)
    {
        var config = new CsvConfiguration(CultureInfo.InvariantCulture)
        {
            HasHeaderRecord = true
        };

        using var sreader = new StreamReader(inFilePath);
        using var reader = new CsvReader(sreader, config);
        using var writer = new StreamWriter(outFilePath);

        var jsonOptions = new JsonSerializerOptions
        {
            Encoder = JavaScriptEncoder.Create(UnicodeRanges.All)
        };
        jsonOptions.Converters.Add(new CustomJsonConverterDateTimeOffset());

        while (reader.Read())
        {
            var inItem = reader.GetRecord<CsvPerson>();
            var outItem = ProcessPerson(inItem);
            writer.WriteLine(JsonSerializer.Serialize(outItem, jsonOptions));
        }
    }

    public Person ProcessPerson(CsvPerson item)
    {
        return new(item.Id, item.Name, item.Age, item.Height, item.IsActive, item.CreatedAt);
    }
}

record struct CsvPerson(
    [property: Index(0)] int Id,
    [property: Index(1)] string Name,
    [property: Index(2)] int Age,
    [property: Index(3)] double Height,
    [property: Index(4)] bool IsActive,
    [property: Index(5)] DateTimeOffset CreatedAt
);

record struct Person(
    [property: JsonPropertyName("id")] int Id,
    [property: JsonPropertyName("name")] string Name,
    [property: JsonPropertyName("age")] int Age,
    [property: JsonPropertyName("height")] double Height,
    [property: JsonPropertyName("is_active")] bool IsActive,
    [property: JsonPropertyName("created_at")] DateTimeOffset CreatedAt
);

class CustomJsonConverterDateTimeOffset : JsonConverter<DateTimeOffset>
{
    public override DateTimeOffset Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    {
        return reader.GetDateTimeOffset();
    }

    public override void Write(Utf8JsonWriter writer, DateTimeOffset value, JsonSerializerOptions options)
    {
        writer.WriteStringValue(value.ToString("yyyy-MM-ddTHH:mm:ssZ"));
    }
}
