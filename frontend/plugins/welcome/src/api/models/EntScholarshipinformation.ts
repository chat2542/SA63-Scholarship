/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntScholarshipinformationEdges,
    EntScholarshipinformationEdgesFromJSON,
    EntScholarshipinformationEdgesFromJSONTyped,
    EntScholarshipinformationEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntScholarshipinformation
 */
export interface EntScholarshipinformation {
    /**
     * ScholarshipName holds the value of the "Scholarship_name" field.
     * @type {string}
     * @memberof EntScholarshipinformation
     */
    scholarshipName?: string;
    /**
     * 
     * @type {EntScholarshipinformationEdges}
     * @memberof EntScholarshipinformation
     */
    edges?: EntScholarshipinformationEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntScholarshipinformation
     */
    id?: number;
}

export function EntScholarshipinformationFromJSON(json: any): EntScholarshipinformation {
    return EntScholarshipinformationFromJSONTyped(json, false);
}

export function EntScholarshipinformationFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntScholarshipinformation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'scholarshipName': !exists(json, 'Scholarship_name') ? undefined : json['Scholarship_name'],
        'edges': !exists(json, 'edges') ? undefined : EntScholarshipinformationEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntScholarshipinformationToJSON(value?: EntScholarshipinformation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Scholarship_name': value.scholarshipName,
        'edges': EntScholarshipinformationEdgesToJSON(value.edges),
        'id': value.id,
    };
}


