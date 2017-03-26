package me.wenzong;

import java.io.IOException;
import java.nio.file.Paths;

import org.apache.lucene.analysis.Analyzer;
import org.apache.lucene.analysis.standard.StandardAnalyzer;
import org.apache.lucene.codecs.simpletext.SimpleTextCodec;
import org.apache.lucene.document.Document;
import org.apache.lucene.document.Field;
import org.apache.lucene.document.FloatPoint;
import org.apache.lucene.document.LongPoint;
import org.apache.lucene.document.StringField;
import org.apache.lucene.document.TextField;
import org.apache.lucene.index.DirectoryReader;
import org.apache.lucene.index.IndexReader;
import org.apache.lucene.index.IndexWriter;
import org.apache.lucene.index.IndexWriterConfig;
import org.apache.lucene.index.IndexWriterConfig.OpenMode;
import org.apache.lucene.queryparser.classic.ParseException;
import org.apache.lucene.queryparser.classic.QueryParser;
import org.apache.lucene.search.IndexSearcher;
import org.apache.lucene.search.Query;
import org.apache.lucene.search.ScoreDoc;
import org.apache.lucene.store.Directory;
import org.apache.lucene.store.FSDirectory;

/**
 * Hello world!
 *
 */
public class App
{
    public static void main(String[] args) throws IOException, ParseException {
        Directory indexDirectory = FSDirectory.open(Paths.get("/tmp/index"));
        //Directory indexDirectory = new RAMDirectory();

        Analyzer analyzer = new StandardAnalyzer();

        IndexWriterConfig iwc = new IndexWriterConfig(analyzer);
        iwc.setOpenMode(OpenMode.CREATE);
        iwc.setCodec(new SimpleTextCodec());
        iwc.setUseCompoundFile(Boolean.FALSE);

        IndexWriter writer = new IndexWriter(indexDirectory, iwc);

        Document doc = new Document();

        Field titleField = new StringField("title", "text", Field.Store.NO);
        doc.add(titleField);

        Field ctxField = new TextField("ctx", "hello world yes sir ok", Field.Store.YES);
        doc.add(ctxField);

        Field stockField = new LongPoint("stock", 2048);
        doc.add(stockField);

        Field priceField = new FloatPoint("price", 1.1F);
        doc.add(priceField);

        writer.addDocument(doc);
        writer.close();

        String field = "title";

        IndexReader reader = DirectoryReader.open(indexDirectory);
        IndexSearcher searcher = new IndexSearcher(reader);

        QueryParser parser = new QueryParser(field, analyzer);
        Query query = parser.parse("text");

        ScoreDoc[] hits = searcher.search(query, 1000).scoreDocs;
        System.out.println(hits.length);

        for (int i = 0; i < hits.length; i++) {
            Document hitDoc = searcher.doc(hits[i].doc);
            System.out.println(hitDoc.get("title"));
        }

        reader.close();
        indexDirectory.close();
    }
}
